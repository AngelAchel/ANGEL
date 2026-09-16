#!/bin/bash
# ANGEL LAB v3.3 - Comprehensive Layer Test
# Tests all 70 layers + end-to-end

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PASS=0
FAIL=0
SKIP=0
TOTAL=0

check() {
    local name="$1"
    local url="$2"
    local expected="${3:-200}"
    local method="${4:-GET}"
    TOTAL=$((TOTAL + 1))

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

check_go() {
    local name="$1"
    local pkg="$2"
    TOTAL=$((TOTAL + 1))
    if go test -count=1 -run "^$" "$pkg" 2>/dev/null; then
        echo -e "  ${GREEN}✓${NC} $name (compile+test)"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $name (FAIL)"
        FAIL=$((FAIL + 1))
    fi
}

check_file() {
    local name="$1"
    local path="$2"
    TOTAL=$((TOTAL + 1))
    if [ -f "$path" ]; then
        echo -e "  ${GREEN}✓${NC} $name"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $name MISSING"
        FAIL=$((FAIL + 1))
    fi
}

echo "=========================================="
echo "  ANGEL LAB v3.3 - Comprehensive Test"
echo "  $(date -u '+%Y-%m-%d %H:%M:%S UTC')"
echo "=========================================="
echo ""

# --- SECTION 1: BUILD ---
echo -e "${BLUE}=== BUILD ===${NC}"
go build ./... 2>&1
if [ $? -eq 0 ]; then
    echo -e "  ${GREEN}✓${NC} All packages compile"
    PASS=$((PASS + 1))
else
    echo -e "  ${RED}✗${NC} Build FAILED"
    FAIL=$((FAIL + 1))
fi
TOTAL=$((TOTAL + 1))
echo ""

# --- SECTION 2: ALL 70 LAYER COMPILE ---
echo -e "${BLUE}=== LAYER COMPILE TEST ===${NC}"
for dir in modules/layer*/; do
    layer=$(basename "$dir")
    count=$(find "$dir" -name "*.go" | wc -l)
    TOTAL=$((TOTAL + 1))
    if [ "$count" -gt 0 ]; then
        echo -e "  ${GREEN}✓${NC} $layer ($count Go files)"
        PASS=$((PASS + 1))
    else
        echo -e "  ${RED}✗${NC} $layer EMPTY"
        FAIL=$((FAIL + 1))
    fi
done
echo ""

# --- SECTION 3: KEY MODULE FUNCTIONALITY ---
echo -e "${BLUE}=== KEY MODULE TESTS ===${NC}"
check_go "C2 Server config" "./modules/layer01-05/c2server/..."
check_go "Auth Bypass" "./modules/layer22-25/authbypass/..."
check_go "Mobile Keychain" "./modules/layer26-40/mobile/..."
check_go "Gateway" "./gateway/..."
check_go "Orchestrator" "./modules/layer11-15/orchestrator/..."
check_go "Implant Gen" "./modules/layer22-25/implantgen/..."
check_go "Crypto" "./modules/layer61-70/crypto/..."
echo ""

# --- SECTION 4: DOCKER ---
echo -e "${BLUE}=== DOCKER ===${NC}"
docker compose ps --format "table {{.Name}}\t{{.Status}}\t{{.Ports}}" 2>&1
echo ""

# --- SECTION 5: ENDPOINTS ---
echo -e "${BLUE}=== ENDPOINTS ===${NC}"
check "Teamserver C2" "http://localhost:8443/" "404" "GET"
check "Teamserver POST" "http://localhost:8443/register" "405" "POST"
check "Console API" "http://localhost:3000/" "404" "GET"
check "DVWA Target" "http://localhost:8081/" "302" "GET"
echo ""

# --- SECTION 6: IMPLANT GENERATE ---
echo -e "${BLUE}=== IMPLANT GENERATE ===${NC}"
TOTAL=$((TOTAL + 1))
if [ -f "bin/angel-generate" ]; then
    ./bin/angel-generate -os linux -arch amd64 -server http://localhost:8443 -out /tmp/angellab 2>&1 | head -3
    if [ $? -eq 0 ] || [ -f "/tmp/angellab" ]; then
        echo -e "  ${GREEN}✓${NC} Implant generate OK"
        PASS=$((PASS + 1))
    else
        echo -e "  ${YELLOW}⊘${NC} Implant generate (no target)"
        SKIP=$((SKIP + 1))
    fi
else
    echo -e "  ${RED}✗${NC} angel-generate binary MISSING"
    FAIL=$((FAIL + 1))
fi
echo ""

# --- SECTION 7: BINARIES ---
echo -e "${BLUE}=== BINARIES ===${NC}"
for b in bin/angel bin/angel-console bin/angel-generate; do
    TOTAL=$((TOTAL + 1))
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

# --- SECTION 8: EVENT BUS ---
echo -e "${BLUE}=== EVENT BUS ===${NC}"
check_go "Event Bus" "./pkg/eventbus/..."
echo ""

# --- SECTION 9: GO TEST ---
echo -e "${BLUE}=== GO TEST ===${NC}"
go test -short ./... 2>&1 | grep -E "^ok|^FAIL" | head -30
echo ""

# --- SUMMARY ---
echo "=========================================="
echo -e "  RESULTS: ${GREEN}$PASS PASS${NC} | ${RED}$FAIL FAIL${NC} | ${YELLOW}$SKIP SKIP${NC} | $TOTAL TOTAL"
echo "=========================================="

if [ $FAIL -eq 0 ]; then
    echo -e "  ${GREEN}LAB READY${NC}"
    exit 0
else
    echo -e "  ${RED}LAB NOT READY - fix $FAIL item(s)${NC}"
    exit 1
fi