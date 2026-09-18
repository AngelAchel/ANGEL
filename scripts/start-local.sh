#!/bin/bash
# ANGEL Start Local Services
# Usage: bash scripts/start-local.sh

set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo "=========================================="
echo "  ANGEL - Start Local Services"
echo "=========================================="
echo ""

# Check env vars - must be set, no fallback to test values
if [ -z "$TEAMSERVER_KEY" ]; then
    echo -e "  ${RED}✗ TEAMSERVER_KEY not set${NC}"
    echo "  Generate: export TEAMSERVER_KEY=\"\$(openssl rand -base64 32)\""
    exit 1
fi
if [ -z "$CRYPTO_KEY" ]; then
    echo -e "  ${RED}✗ CRYPTO_KEY not set${NC}"
    echo "  Generate: export CRYPTO_KEY=\"\$(openssl rand -base64 32)\""
    exit 1
fi
if [ -z "$JWT_SECRET" ]; then
    echo -e "  ${RED}✗ JWT_SECRET not set${NC}"
    echo "  Generate: export JWT_SECRET=\"\$(openssl rand -base64 32)\""
    exit 1
fi

# Create lab directories
mkdir -p lab/data lab/logs lab/implants lab/configs

# Start teamserver
echo "--- Starting Teamserver ---"
./bin/angel -bind 127.0.0.1 -port 8443 > lab/logs/teamserver.log 2>&1 &
TEAMSERVER_PID=$!
echo $TEAMSERVER_PID > lab/logs/teamserver.pid
sleep 1
if kill -0 $TEAMSERVER_PID 2>/dev/null; then
    echo -e "  ${GREEN}✓${NC} Teamserver started (PID: $TEAMSERVER_PID)"
else
    echo -e "  ${RED}✗${NC} Teamserver failed to start"
fi

# Start console
echo "--- Starting Console ---"
./bin/angel-console -addr 127.0.0.1 -port 3000 > lab/logs/console.log 2>&1 &
CONSOLE_PID=$!
echo $CONSOLE_PID > lab/logs/console.pid
sleep 1
if kill -0 $CONSOLE_PID 2>/dev/null; then
    echo -e "  ${GREEN}✓${NC} Console started (PID: $CONSOLE_PID)"
else
    echo -e "  ${RED}✗${NC} Console failed to start"
fi

# Start rules loader
echo "--- Starting Rules Loader ---"
./bin/angel-rules > lab/logs/rules.log 2>&1 &
RULES_PID=$!
echo $RULES_PID > lab/logs/rules.pid
sleep 1
if kill -0 $RULES_PID 2>/dev/null; then
    echo -e "  ${GREEN}✓${NC} Rules loader started (PID: $RULES_PID)"
else
    echo -e "  ${RED}✗${NC} Rules loader failed to start"
fi

echo ""
echo "=========================================="
echo "  Services started"
echo "=========================================="
echo "  Teamserver:  http://127.0.0.1:8443"
echo "  Console:     http://127.0.0.1:3000"
echo "  Rules:       http://127.0.0.1:9444"
echo ""
echo "  PIDs saved to lab/logs/*.pid"
echo "  Logs saved to lab/logs/*.log"
echo "  Run: bash scripts/health-check.sh"
echo ""