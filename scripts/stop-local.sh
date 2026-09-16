#!/bin/bash
# ANGEL Stop Local Services
# Usage: bash scripts/stop-local.sh

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "=========================================="
echo "  ANGEL - Stop Local Services"
echo "=========================================="
echo ""

# Stop by PID files
for service in teamserver console rules; do
    PID_FILE="lab/logs/${service}.pid"
    if [ -f "$PID_FILE" ]; then
        PID=$(cat "$PID_FILE")
        if kill -0 "$PID" 2>/dev/null; then
            kill "$PID" 2>/dev/null
            echo -e "  ${GREEN}✓${NC} Stopped $service (PID: $PID)"
        else
            echo -e "  ${YELLOW}⚠${NC} $service not running (stale PID: $PID)"
        fi
        rm -f "$PID_FILE"
    else
        echo -e "  ${YELLOW}⚠${NC} No PID file for $service"
    fi
done

# Stop by process name (fallback)
pkill -f "angel-cgo0" 2>/dev/null && echo -e "  ${GREEN}✓${NC} Killed angel processes" || true
pkill -f "angel-console-cgo0" 2>/dev/null && true || true
pkill -f "angel-rules-cgo0" 2>/dev/null && true || true

echo ""
echo "  All services stopped"
echo ""