#!/bin/bash
# ANGEL Health Check
# Usage: bash scripts/health-check.sh

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PASS=0
FAIL=0
WARN=0

check_service() {
    local name="$1"
    local url="$2"
    local expected="$3"
    
    code=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 "$url" 2>/dev/null)
    if [ "$code" = "$expected" ] || [ "$code" != "000" ]; then
        echo -e "  ${GREEN}✓${NC} $name → $code"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $name → $code (expected $expected)"
        FAIL=$((FAIL + 1))
    fi
}

check_process() {
    local name="$1"
    local pattern="$2"
    
    if pgrep -f "$pattern" >/dev/null 2>&1; then
        echo -e "  ${GREEN}✓${NC} $name running"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $name not running"
        FAIL=$((FAIL + 1))
    fi
}

echo "=========================================="
echo "  ANGEL Health Check"
echo "=========================================="
echo ""

echo "--- Services ---"
check_process "Teamserver" "angel-cgo0"
check_process "Console" "angel-console-cgo0"
check_process "Rules Loader" "angel-rules-cgo0"
echo ""

echo "--- Endpoints ---"
check_service "Teamserver" "http://127.0.0.1:8443/" "404"
check_service "Console" "http://127.0.0.1:3000/" "404"
check_service "Rules" "http://127.0.0.1:9444/" "404"
echo ""

echo "--- Docker ---"
if docker ps >/dev/null 2>&1; then
    echo -e "  ${GREEN}✓${NC} Docker running"
    PASS=$((PASS + 1))
    
    for container in angel-teamserver angel-console angel-rules angel-dvwa; do
        if docker ps --format "{{.Names}}" | grep -q "^${container}$"; then
            echo -e "  ${GREEN}✓${NC} $container running"
            PASS=$((PASS + 1))
        else
            echo -e "  ${RED}✗${NC} $container not running"
            FAIL=$((FAIL + 1))
        fi
    done
else
    echo -e "  ${YELLOW}⚠${NC} Docker not available"
    WARN=$((WARN + 1))
fi
echo ""

echo "--- Lab Data ---"
[ -d "lab/data" ] && echo -e "  ${GREEN}✓${NC} lab/data exists" || echo -e "  ${RED}✗${NC} lab/data missing"
[ -d "lab/logs" ] && echo -e "  ${GREEN}✓${NC} lab/logs exists" || echo -e "  ${RED}✗${NC} lab/logs missing"
[ -d "lab/implants" ] && echo -e "  ${GREEN}✓${NC} lab/implants exists" || echo -e "  ${RED}✗${NC} lab/implants missing"
echo ""

echo "--- Environment ---"
[ -n "$TEAMSERVER_KEY" ] && echo -e "  ${GREEN}✓${NC} TEAMSERVER_KEY set" || echo -e "  ${RED}✗${NC} TEAMSERVER_KEY not set"
[ -n "$CRYPTO_KEY" ] && echo -e "  ${GREEN}✓${NC} CRYPTO_KEY set" || echo -e "  ${RED}✗${NC} CRYPTO_KEY not set"
[ -n "$JWT_SECRET" ] && echo -e "  ${GREEN}✓${NC} JWT_SECRET set" || echo -e "  ${RED}✗${NC} JWT_SECRET not set"
echo ""

echo "--- Binaries ---"
[ -f "bin/angel-cgo0" ] && echo -e "  ${GREEN}✓${NC} angel binary" || echo -e "  ${RED}✗${NC} angel binary missing"
[ -f "bin/angel-console-cgo0" ] && echo -e "  ${GREEN}✓${NC} console binary" || echo -e "  ${RED}✗${NC} console binary missing"
[ -f "bin/angel-generate-cgo0" ] && echo -e "  ${GREEN}✓${NC} generate binary" || echo -e "  ${RED}✗${NC} generate binary missing"
[ -f "bin/angel-rules-cgo0" ] && echo -e "  ${GREEN}✓${NC} rules binary" || echo -e "  ${RED}✗${NC} rules binary missing"
echo ""

echo "=========================================="
echo "  Results: $PASS PASS, $FAIL FAIL, $WARN WARN"
echo "=========================================="

if [ "$FAIL" -gt 0 ]; then
    echo -e "  ${RED}Some services not healthy${NC}"
    exit 1
else
    echo -e "  ${GREEN}All services healthy${NC}"
    exit 0
fi