#!/bin/bash
# ANGEL Monitoring Script
# Usage: bash scripts/monitor.sh <domain>

DOMAIN="$1"

if [ -z "$DOMAIN" ]; then
    echo "Usage: bash scripts/monitor.sh <domain>"
    echo "Example: bash scripts/monitor.sh angel.example.com"
    exit 1
fi

echo "=== ANGEL Monitor ==="
echo "Target: $DOMAIN"
echo "Time: $(date)"
echo ""

# Check services
echo "=== SERVICES ==="
for service in angel-teamserver angel-console angel-rules; do
    if systemctl is-active --quiet "$service"; then
        echo "  ✓ $service: RUNNING"
    else
        echo "  ✗ $service: DOWN"
    fi
done

# Check ports
echo ""
echo "=== PORTS ==="
for port in 443 8443 3000 9444; do
    if ss -tlnp | grep -q ":$port "; then
        echo "  ✓ Port $port: LISTENING"
    else
        echo "  ✗ Port $port: NOT LISTENING"
    fi
done

# Check TLS certificate
echo ""
echo "=== TLS ==="
cert_expiry=$(echo | openssl s_client -connect "$DOMAIN:443" -servername "$DOMAIN" 2>/dev/null | openssl x509 -enddate -noout 2>/dev/null | cut -d= -f2)
if [ -n "$cert_expiry" ]; then
    echo "  Certificate expiry: $cert_expiry"
else
    echo "  ✗ Cannot connect to TLS"
fi

# Check disk space
echo ""
echo "=== DISK ==="
df -h /opt/ANGEL 2>/dev/null | tail -1 | awk '{print "  Used: "$3" Free: "$4" ("$5"%)"}'

# Check memory
echo ""
echo "=== MEMORY ==="
free -h | grep -i mem | awk '{print "  Total: "$2" Used: "$3" Free: "$4}'

# Check CPU
echo ""
echo "=== CPU ==="
uptime | awk -F'load average:' '{print "  Load: "$2}'

# Check logs
echo ""
echo "=== RECENT LOGS ==="
if [ -d "lab/logs" ]; then
    echo "Last 5 lines from teamserver.log:"
    tail -5 lab/logs/teamserver.log 2>/dev/null || echo "  No logs"
    echo ""
    echo "Last 5 lines from console.log:"
    tail -5 lab/logs/console.log 2>/dev/null || echo "  No logs"
else
    echo "  No logs directory"
fi

echo ""
echo "=== HEALTH CHECK ==="
bash scripts/health-check-prod.sh "$DOMAIN" 2>/dev/null || echo "  Health check skipped"