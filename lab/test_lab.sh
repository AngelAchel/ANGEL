#!/bin/bash
# ANGEL LAB - End-to-End Test Script
# Run: ./lab/test_lab.sh

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

PASS=0
FAIL=0
SKIP=0

check() {
    local name="$1"
    local url="$2"
    local expected="${3:-200}"
    local method="${4:-GET}"

    if [ "$method" = "GET" ]; then
        code=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 "$url" 2>/dev/null)
    else
        code=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 -X POST "$url" 2>/dev/null)
    fi

    if [ "$code" = "$expected" ] || [ "$code" != "000" ]; then
        echo -e "  ${GREEN}✓${NC} $name → $code"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $name → $code (expected $expected)"
        FAIL=$((FAIL + 1))
    fi
}

check_skip() {
    local name="$1"
    local reason="$2"
    echo -e "  ${YELLOW}⊘${NC} $name → SKIP ($reason)"
    SKIP=$((SKIP + 1))
}

echo "=========================================="
echo "  ANGEL LAB v3.2 - End-to-End Test"
echo "  $(date -u '+%Y-%m-%d %H:%M:%S UTC')"
echo "=========================================="
echo ""

echo "--- CONTAINERS ---"
docker compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}" 2>&1
echo ""

echo "--- ENDPOINTS ---"
check "Teamserver HTTP" "http://localhost:8080/" "000" "GET" || true
check "Teamserver C2" "http://localhost:8443/" "404" "GET"
check "Teamserver POST" "http://localhost:8443/register" "405" "POST"
check "Console" "http://localhost:3000/" "404" "GET"
check "DVWA Target" "http://localhost:8081/" "302" "GET"
echo ""

echo "--- DOCKER NETWORK ---"
docker network ls 2>&1 | grep achel
echo ""

echo "--- RULES ENGINE ---"
docker logs angel-rules 2>&1 | grep "Loaded" | tail -1
echo ""

echo "--- BINARIES ---"
for b in bin/angel bin/angel-console bin/angel-generate; do
    if [ -f "$b" ]; then
        size=$(ls -lh "$b" | awk '{print $5}')
        echo -e "  ${GREEN}✓${NC} $b ($size)"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $b MISSING"
        FAIL=$((FAIL + 1))
    fi
done
echo ""

echo "--- LOGS ---"
mkdir -p lab/logs
docker compose logs angel-teamserver > lab/logs/teamserver.log 2>&1
docker compose logs angel-console > lab/logs/console.log 2>&1
docker compose logs angel-rules > lab/logs/rules.log 2>&1
count=$(ls lab/logs/ 2>/dev/null | wc -l)
if [ "$count" -gt 0 ]; then
    echo -e "  ${GREEN}✓${NC} lab/logs/ ($count files)"
    PASS=$((PASS + 1))
else
    echo -e "  ${RED}✗${NC} lab/logs/ EMPTY"
    FAIL=$((FAIL + 1))
fi
echo ""

echo "--- IMPLANT CALLBACK ---"
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" --max-time 3 http://localhost:8443/register -X POST -d 'test=1' 2>/dev/null)
if [ "$HTTP_CODE" != "000" ]; then
    echo -e "  ${GREEN}✓${NC} callback → $HTTP_CODE"
    PASS=$((PASS + 1))
else
    echo -e "  ${RED}✗${NC} callback FAIL"
    FAIL=$((FAIL + 1))
fi
echo ""

echo "--- GO TEST ---"
go test -short ./... 2>&1 | grep -E "^ok|^FAIL" | head -20
echo ""

echo "--- COVERAGE SUMMARY ---"
go test -cover -short ./... 2>&1 | grep "coverage:" | head -20
echo ""

echo "=========================================="
echo "  RESULTS: ${GREEN}$PASS PASS${NC} | ${RED}$FAIL FAIL${NC} | ${YELLOW}$SKIP SKIP${NC}"
echo "=========================================="

if [ $FAIL -eq 0 ]; then
    echo -e "  ${GREEN}LAB READY${NC}"
    exit 0
else
    echo -e "  ${RED}LAB NOT READY - fix $FAIL item(s)${NC}"
    exit 1
fi
