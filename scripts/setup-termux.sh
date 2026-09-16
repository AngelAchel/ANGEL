#!/bin/bash
# ANGEL Setup for Termux
# Usage: bash scripts/setup-termux.sh
#
# WARNING: ANGEL is NOT supported on Termux.
# This script attempts to set up ANGEL on Termux but may not work.
# Termux lacks: Go, Docker, make, gcc, nmap, iptables, root access.
# Cross-compile only - runtime NOT tested.

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo "=========================================="
echo "  ANGEL Setup - Termux (UNSUPPORTED)"
echo "=========================================="
echo ""
echo -e "  ${RED}WARNING: ANGEL is NOT supported on Termux${NC}"
echo -e "  ${RED}This script will attempt setup but may fail${NC}"
echo ""

# Check if Termux
if [ -z "$PREFIX" ]; then
    echo -e "  ${RED}✗ Not a Termux environment${NC}"
    exit 1
fi

echo "--- Termux environment detected ---"
echo "  PREFIX: $PREFIX"
echo "  Architecture: $(uname -m)"
echo ""

# Try to install Go (may not be available)
echo "--- Attempting Go installation ---"
pkg install -y golang 2>/dev/null && echo -e "  ${GREEN}✓${NC} Go installed" || {
    echo -e "  ${RED}✗ Go NOT available in Termux${NC}"
    echo -e "  ${YELLOW}Manual: compile Go for Android or use static binary${NC}"
}

# Try to install dependencies
echo "--- Attempting dependency installation ---"
pkg install -y git curl python 2>/dev/null
echo -e "  ${YELLOW}⚠${NC} git, curl, python may be installed"

# Docker is NOT available on Termux
echo -e "  ${RED}✗ Docker NOT available on Termux${NC}"
echo -e "  ${YELLOW}Docker services (DVWA, rules) will NOT work${NC}"

# Try to cross-compile (may fail)
echo "--- Attempting cross-compile ---"
cd /data/data/com.termux/files/home 2>/dev/null || exit 1
if command -v go >/dev/null 2>&1; then
    CGO_ENABLED=0 GOOS=android GOARCH=arm64 go build -o angel-termux ./cmd/teamserver 2>/dev/null && \
    echo -e "  ${GREEN}✓${NC} Cross-compile OK" || \
    echo -e "  ${RED}✗ Cross-compile FAILED${NC}"
else
    echo -e "  ${RED}✗ Go not available - cannot cross-compile${NC}"
fi

# Create lab directories
mkdir -p ~/storage/shared/angel-lab/data ~/storage/shared/angel-lab/logs 2>/dev/null

echo ""
echo "=========================================="
echo "  Status: NOT SUPPORTED"
echo "=========================================="
echo ""
echo "Termux limitations:"
echo "  - No Go package"
echo "  - No Docker"
echo "  - No make"
echo "  - No gcc"
echo "  - No iptables"
echo "  - No root access"
echo "  - Limited storage"
echo "  - Limited network binding"
echo ""
echo "For native ANGEL, use Kali Linux or Ubuntu."
echo ""