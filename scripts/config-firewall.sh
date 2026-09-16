#!/bin/bash
# ANGEL Firewall Configuration
# Run on production VPS
# Usage: bash scripts/config-firewall.sh

set -e

echo "Configuring firewall..."

# Reset firewall
ufw --force reset

# Default policy
ufw default deny incoming
ufw default allow outgoing

# Allow SSH
ufw allow 22/tcp
ufw allow 22/udp

# Allow ANGEL ports
ufw allow 443/tcp    # HTTPS
ufw allow 8443/tcp   # Teamserver API
ufw allow 3000/tcp   # Console API
ufw allow 9444/tcp   # Rules engine

# Allow DNS
ufw allow 53/tcp
ufw allow 53/udp

# Allow SMB
ufw allow 445/tcp
ufw allow 445/udp

# Allow DNS tunnel port
ufw allow 8444/tcp

# Enable firewall
ufw --force enable

echo ""
echo "Firewall configured:"
ufw status numbered

echo ""
echo "Open ports:"
echo "  22/tcp   - SSH"
echo "  443/tcp  - HTTPS"
echo "  8443/tcp - Teamserver API"
echo "  3000/tcp - Console API"
echo "  9444/tcp - Rules engine"
echo "  53/tcp   - DNS"
echo "  445/tcp  - SMB"
echo "  8444/tcp - DNS tunnel"