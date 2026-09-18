#!/bin/bash
# ANGEL Production Deployment Script
# Run on fresh VPS Kali/Ubuntu 22.04+
# Usage: bash scripts/deploy-production.sh <domain> <vps_ip>

set -e

DOMAIN="$1"
VPS_IP="$2"

if [ -z "$DOMAIN" ] || [ -z "$VPS_IP" ]; then
    echo "Usage: bash scripts/deploy-production.sh <domain> <vps_ip>"
    echo "Example: bash scripts/deploy-production.sh angel.example.com 192.168.1.100"
    exit 1
fi

echo "=========================================="
echo "  ANGEL Production Deployment"
echo "  Domain: $DOMAIN"
echo "  VPS IP: $VPS_IP"
echo "=========================================="

# Step 1: Update system
echo "[1/6] Updating system..."
apt update && apt upgrade -y
apt install -y curl wget git nginx ufw fail2ban

# Step 2: Install Go
echo "[2/6] Installing Go..."
if ! command -v go &> /dev/null; then
    wget -q https://go.dev/dl/go1.27.0.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.27.0.linux-amd64.tar.gz
    rm go1.27.0.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    export PATH=$PATH:/usr/local/go/bin
fi
go version

# Step 3: Clone repo
echo "[3/6] Cloning ANGEL repo..."
cd /opt
if [ -d "ANGEL" ]; then
    rm -rf ANGEL
fi
git clone https://github.com/AngelAchel/ANGEL.git
cd ANGEL

# Step 4: Build
echo "[4/6] Building..."
export TEAMSERVER_KEY="$(openssl rand -base64 32)"
export CRYPTO_KEY="$(openssl rand -base64 32)"
export JWT_SECRET="$(openssl rand -base64 32)"
export DATABASE_ENCRYPTION_KEY="$(openssl rand -base64 32)"

make build

# Save env vars for systemd
cat > /etc/angel.env << EOF
TEAMSERVER_KEY=$TEAMSERVER_KEY
CRYPTO_KEY=$CRYPTO_KEY
JWT_SECRET=$JWT_SECRET
DATABASE_ENCRYPTION_KEY=$DATABASE_ENCRYPTION_KEY
DOMAIN=$DOMAIN
VPS_IP=$VPS_IP
EOF

echo "Env vars saved to /etc/angel.env (permissions: 600)"
chmod 600 /etc/angel.env
echo ""
echo "IMPORTANT: Generated secrets are stored in /etc/angel.env"
echo "          View with: sudo cat /etc/angel.env"
echo "          Rotate anytime with: openssl rand -base64 32"

# Step 5: TLS cert (Let's Encrypt)
echo "[5/6] Setting up TLS..."
apt install -y certbot python3-certbot-nginx

# Configure nginx first (temporary HTTP for cert)
cat > /etc/nginx/sites-available/angel << 'NGINX'
server {
    listen 80;
    server_name _DOMAIN_;
    location /.well-known/acme-challenge/ {
        root /var/www/certbot;
    }
    location / {
        return 301 https://$host$request_uri;
    }
}
NGINX

sed -i "s/_DOMAIN_/$DOMAIN/g" /etc/nginx/sites-available/angel
ln -sf /etc/nginx/sites-available/angel /etc/nginx/sites-enabled/
rm -f /etc/nginx/sites-enabled/default
nginx -t && systemctl restart nginx

# Get cert
mkdir -p /var/www/certbot
certbot certonly --standalone -d "$DOMAIN" --non-interactive --agree-tos --register-unsafely-without-email || true

# Step 6: Systemd services
echo "[6/6] Creating systemd services..."

cat > /etc/systemd/system/angel-teamserver.service << 'EOF'
[Unit]
Description=ANGEL Teamserver
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/ANGEL
EnvironmentFile=/etc/angel.env
ExecStart=/opt/ANGEL/bin/angel -bind 0.0.0.0 -port 8443
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

cat > /etc/systemd/system/angel-console.service << 'EOF'
[Unit]
Description=ANGEL Console
After=network.target angel-teamserver.service

[Service]
Type=simple
User=root
WorkingDirectory=/opt/ANGEL
EnvironmentFile=/etc/angel.env
ExecStart=/opt/ANGEL/bin/angel-console -addr 0.0.0.0 -port 3000
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

cat > /etc/systemd/system/angel-rules.service << 'EOF'
[Unit]
Description=ANGEL Rules Engine
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/ANGEL
EnvironmentFile=/etc/angel.env
ExecStart=/opt/ANGEL/bin/angel-rules -bind 0.0.0.0 -port 9444
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable angel-teamserver angel-console angel-rules
systemctl start angel-teamserver angel-console angel-rules

echo ""
echo "=========================================="
echo "  DEPLOYMENT COMPLETE"
echo "=========================================="
echo "Teamserver: https://$DOMAIN:8443"
echo "Console:    https://$DOMAIN:3000"
echo "Rules:      https://$DOMAIN:9444"
echo ""
echo "Next steps:"
echo "1. Configure nginx reverse proxy with TLS"
echo "2. Open firewall ports: 443, 8443, 3000"
echo "3. Test: curl https://$DOMAIN:8443"
echo "4. View secrets: sudo cat /etc/angel.env"