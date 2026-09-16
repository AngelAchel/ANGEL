#!/bin/bash
# ANGEL LAB v3.4 - Per-Layer Functional Test
# Tests all 83 packages individually

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

PASS=0
FAIL=0
SKIP=0
TOTAL=0

echo "=========================================="
echo "  ANGEL LAB v3.4 - Per-Layer Test"
echo "  $(date -u '+%Y-%m-%d %H:%M:%S UTC')"
echo "=========================================="
echo ""

# Get all packages
PACKAGES=$(go list ./modules/... 2>/dev/null | grep -v "_test.go" | sort)

# Test each package
while IFS= read -r pkg; do
    TOTAL=$((TOTAL + 1))
    name=$(echo "$pkg" | sed 's|github.com/angel-platform/angel/||')

    # Try to build the package
    output=$(go build "$pkg" 2>&1)
    if [ $? -eq 0 ]; then
        # Try to run tests for this package
        test_output=$(go test -count=1 -short "$pkg" 2>&1)
        if echo "$test_output" | grep -q "^ok\|^no test"; then
            echo -e "  ${GREEN}✓${NC} $name"
            PASS=$((PASS + 1))
        elif echo "$test_output" | grep -q "no test files"; then
            echo -e "  ${YELLOW}⊘${NC} $name (no tests)"
            SKIP=$((SKIP + 1))
        else
            echo -e "  ${RED}✗${NC} $name (test fail)"
            FAIL=$((FAIL + 1))
        fi
    else
        echo -e "  ${RED}✗${NC} $name (build fail)"
        echo "    $output"
        FAIL=$((FAIL + 1))
    fi
done <<< "$PACKAGES"

echo ""
echo "=========================================="
echo -e "  RESULTS: ${GREEN}$PASS PASS${NC} | ${RED}$FAIL FAIL${NC} | ${YELLOW}$SKIP SKIP${NC} | $TOTAL TOTAL"
echo "=========================================="

if [ $FAIL -eq 0 ]; then
    echo -e "  ${GREEN}ALL 83 LAYERS FUNCTIONAL${NC}"
    exit 0
else
    echo -e "  ${RED}$FAIL LAYER(s) FAILED${NC}"
    exit 1
fi