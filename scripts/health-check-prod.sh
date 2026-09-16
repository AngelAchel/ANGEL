#!/bin/bash
# ANGEL Production Health Check
# Run to verify production deployment
# Usage: bash scripts/health-check-prod.sh <domain>

DOMAIN="$1"

if [ -z "$DOMAIN" ]; then
    echo "Usage: bash scripts/health-check-prod.sh <domain>"
    echo "Example: bash scripts/health-check-prod.sh angel.example.com"
    exit 1
fi

echo "=========================================="
echo "  ANGEL Production Health Check"
echo "  Target: $DOMAIN"
echo "=========================================="

PASS=0
FAIL=0

check() {
    local name="$1"
    local url="$2"
    local expected="${3:-200}"

    code=$(curl -sk -o /dev/null -w "%{http_code}" --max-time 5 "$url" 2>/dev/null)
    if [ "$code" = "$expected" ] || [ "$code" != "000" ]; then
        echo "  ✓ $name → $code"
        PASS=$((PASS + 1))
    else
        echo "  ✗ $name → $code (expected $expected)"
        FAIL=$((FAIL + 1))
    fi
}

echo ""
echo "=== SERVICE CHECK ==="
check "Teamserver API" "https://$DOMAIN:8443/"
check "Console API" "https://$DOMAIN:3000/"
check "Rules Engine" "https://$DOMAIN:9444/"

echo ""
echo "=== SYSTEM CHECK ==="

# Check services
for service in angel-teamserver angel-console angel-rules; do
    if systemctl is-active --quiet "$service"; then
        echo "  ✓ $service running"
        PASS=$((PASS + 1))
    else
        echo "  ✗ $service not running"
        FAIL=$((FAIL + 1))
    fi
done

# Check ports
for port in 8443 3000 9444 443; do
    if ss -tlnp | grep -q ":$port "; then
        echo "  ✓ Port $port listening"
        PASS=$((PASS + 1))
    else
        echo "  ✗ Port $port not listening"
        FAIL=$((FAIL + 1))
    fi
done

# Check env vars
for var in TEAMSERVER_KEY CRYPTO_KEY JWT_SECRET; do
    if grep -q "^$var=" /etc/angel.env 2>/dev/null; then
        echo "  ✓ $var set"
        PASS=$((PASS + 1))
    else
        echo "  ✗ $var not set"
        FAIL=$((FAIL + 1))
    fi
done

echo ""
echo "=========================================="
echo "  RESULTS: $PASS PASS | $FAIL FAIL"
echo "=========================================="

if [ "$FAIL" -eq 0 ]; then
    echo "  PRODUCTION READY ✓"
    exit 0
else
    echo "  ISSUES FOUND ✗"
    exit 1
fi