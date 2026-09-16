#!/bin/bash
# ANGEL Setup for Kali Linux
# Usage: bash scripts/setup-kali.sh

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo "=========================================="
echo "  ANGEL Setup - Kali Linux"
echo "=========================================="
echo ""

# Update system
echo "--- Updating system ---"
sudo apt update -qq
echo -e "  ${GREEN}✓${NC} System updated"

# Install mandatory dependencies
echo "--- Installing mandatory dependencies ---"
sudo apt install -y -qq golang-go git make curl gcc 2>/dev/null
echo -e "  ${GREEN}✓${NC} Go, git, make, curl, gcc installed"

# Install optional dependencies
echo "--- Installing optional dependencies ---"
sudo apt install -y -qq nmap netcat-openbsd python3 iptables 2>/dev/null
echo -e "  ${GREEN}✓${NC} nmap, netcat, python3, iptables installed"

# Install Docker
echo "--- Installing Docker ---"
sudo apt install -y -qq docker.io docker-compose-v2 2>/dev/null
sudo usermod -aG docker $USER 2>/dev/null || true
echo -e "  ${GREEN}✓${NC} Docker installed"

# Clone ANGEL
echo "--- Cloning ANGEL ---"
if [ ! -d "ANGEL" ]; then
    git clone https://github.com/AngelAchel/ANGEL.git
    cd ANGEL
else
    cd ANGEL
    git pull
fi
echo -e "  ${GREEN}✓${NC} ANGEL cloned"

# Build
echo "--- Building ANGEL ---"
CGO_ENABLED=0 make build
echo -e "  ${GREEN}✓${NC} ANGEL built"

# Set environment variables
echo "--- Environment Variables ---"
cat >> ~/.bashrc << 'EOF'
export TEAMSERVER_KEY="change-me"
export CRYPTO_KEY="change-me"
export JWT_SECRET="change-me"
EOF
echo -e "  ${GREEN}✓${NC} Env vars added to ~/.bashrc"
echo -e "  ${YELLOW}⚠${NC} Edit ~/.bashrc with real values"

# Create lab directories
echo "--- Creating lab directories ---"
mkdir -p lab/data lab/logs lab/implants lab/configs
echo -e "  ${GREEN}✓${NC} Lab directories created"

echo ""
echo "=========================================="
echo "  Setup complete!"
echo "=========================================="
echo ""
echo "Next steps:"
echo "  1. Edit ~/.bashrc with real credentials"
echo "  2. source ~/.bashrc"
echo "  3. bash scripts/start-local.sh"
echo "  4. bash scripts/health-check.sh"
echo ""