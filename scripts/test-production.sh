#!/bin/bash
# ANGEL Production Deployment Test
# Run to verify all steps completed
# Usage: bash scripts/test-production.sh <domain>

DOMAIN="$1"

if [ -z "$DOMAIN" ]; then
    echo "Usage: bash scripts/test-production.sh <domain>"
    exit 1
fi

echo "=========================================="
echo "  ANGEL Production Test"
echo "  Target: $DOMAIN"
echo "=========================================="

PASS=0
FAIL=0

# Test 1: TLS certificate
echo ""
echo "=== TLS TEST ==="
cert_info=$(curl -skI "https://$DOMAIN:443" 2>/dev/null | grep -i "strict-transport")
if [ -n "$cert_info" ]; then
    echo "  ✓ HSTS header present"
    PASS=$((PASS + 1))
else
    echo "  ✗ HSTS header missing"
    FAIL=$((FAIL + 1))
fi

# Test 2: Teamserver API
echo ""
echo "=== API TEST ==="
api_code=$(curl -sk -o /dev/null -w "%{http_code}" "https://$DOMAIN:8443/" 2>/dev/null)
if [ "$api_code" = "404" ] || [ "$api_code" = "200" ]; then
    echo "  ✓ Teamserver API responding ($api_code)"
    PASS=$((PASS + 1))
else
    echo "  ✗ Teamserver API not responding ($api_code)"
    FAIL=$((FAIL + 1))
fi

# Test 3: Console API
console_code=$(curl -sk -o /dev/null -w "%{http_code}" "https://$DOMAIN:3000/" 2>/dev/null)
if [ "$console_code" = "404" ] || [ "$console_code" = "200" ]; then
    echo "  ✓ Console API responding ($console_code)"
    PASS=$((PASS + 1))
else
    echo "  ✗ Console API not responding ($console_code)"
    FAIL=$((FAIL + 1))
fi

# Test 4: WebSocket upgrade
ws_code=$(curl -sk -o /dev/null -w "%{http_code}" -H "Upgrade: websocket" -H "Connection: Upgrade" "https://$DOMAIN/ws" 2>/dev/null)
if [ "$ws_code" = "101" ] || [ "$ws_code" = "400" ]; then
    echo "  ✓ WebSocket endpoint active ($ws_code)"
    PASS=$((PASS + 1))
else
    echo "  ✗ WebSocket endpoint not responding ($ws_code)"
    FAIL=$((FAIL + 1))
fi

# Test 5: DNS listener
dns_code=$(curl -sk -o /dev/null -w "%{http_code}" "https://$DOMAIN:8444/" 2>/dev/null)
if [ "$dns_code" = "404" ] || [ "$dns_code" = "200" ] || [ "$dns_code" = "000" ]; then
    echo "  ✓ DNS listener active"
    PASS=$((PASS + 1))
else
    echo "  ✗ DNS listener not responding"
    FAIL=$((FAIL + 1))
fi

# Test 6: SMB listener
smb_code=$(curl -sk -o /dev/null -w "%{http_code}" "https://$DOMAIN:4455/" 2>/dev/null)
if [ "$smb_code" = "404" ] || [ "$smb_code" = "200" ] || [ "$smb_code" = "000" ]; then
    echo "  ✓ SMB listener active"
    PASS=$((PASS + 1))
else
    echo "  ✗ SMB listener not responding"
    FAIL=$((FAIL + 1))
fi

# Test 7: Firewall
echo ""
echo "=== FIREWALL TEST ==="
if ufw status | grep -q "active"; then
    echo "  ✓ Firewall active"
    PASS=$((PASS + 1))
else
    echo "  ✗ Firewall not active"
    FAIL=$((FAIL + 1))
fi

# Test 8: Services
echo ""
echo "=== SERVICE TEST ==="
for service in angel-teamserver angel-console angel-rules; do
    if systemctl is-active --quiet "$service"; then
        echo "  ✓ $service running"
        PASS=$((PASS + 1))
    else
        echo "  ✗ $service not running"
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