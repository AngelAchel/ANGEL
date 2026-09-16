# ANGEL Production Deployment Guide

## Quick Start

```bash
# 1. Deploy to VPS
bash scripts/deploy-production.sh <domain> <vps_ip>

# 2. Configure nginx
bash scripts/config-nginx.sh <domain>

# 3. Configure firewall
bash scripts/config-firewall.sh

# 4. Health check
bash scripts/health-check-prod.sh <domain>

# 5. Full test
bash scripts/test-production.sh <domain>
```

## Architecture

```
Internet
  │
  ▼
┌─────────────┐
│  Nginx      │  HTTPS (443) + Reverse Proxy
│  (TLS)      │
└────┬────────┘
     │
     ├──► Teamserver (8443) - C2 API
     ├──► Console (3000) - API Gateway
     └──► Rules (9444) - Rules Engine
```

## Ports

| Port | Service | Protocol |
|------|---------|----------|
| 443 | HTTPS (nginx) | TCP |
| 8443 | Teamserver API | TCP |
| 3000 | Console API | TCP |
| 9444 | Rules Engine | TCP |
| 8444 | DNS Listener | TCP/UDP |
| 4455 | SMB Listener | TCP |
| 53 | DNS | UDP |

## Environment Variables

Set in `/etc/angel.env`:

```bash
TEAMSERVER_KEY=<strong-random-key-32+char>
CRYPTO_KEY=<strong-random-key-32+char>
JWT_SECRET=<strong-random-secret>
DATABASE_ENCRYPTION_KEY=<strong-random-key>
```

## Services

```bash
# Check status
systemctl status angel-teamserver
systemctl status angel-console
systemctl status angel-rules

# Restart
systemctl restart angel-teamserver
systemctl restart angel-console
systemctl restart angel-rules

# Logs
journalctl -u angel-teamserver -f
journalctl -u angel-console -f
journalctl -u angel-rules -f
```

## Client Usage

```bash
# Connect to teamserver
./angel-console -server https://<domain>:8443 -jwt <JWT_SECRET>

# Generate implant
./angel-generate -os linux -arch amd64 -server https://<domain>:8443

# Generate Windows implant
./angel-generate -os windows -arch amd64 -server https://<domain>:8443
```

## SSL Certificate

Let's Encrypt (auto-renew):
```bash
certbot certonly --standalone -d <domain>
certbot renew --dry-run  # Test renewal
```

## Backup

```bash
# Backup lab data
tar -czf angel-backup-$(date +%Y%m%d).tar.gz lab/data/ lab/logs/ lab/implants/

# Restore
tar -xzf angel-backup-YYYYMMDD.tar.gz -C /opt/ANGEL/
```

## Monitoring

```bash
# Health check
bash scripts/health-check-prod.sh <domain>

# Full test
bash scripts/test-production.sh <domain>

# Check logs
docker compose logs -f angel-teamserver
```

## Troubleshooting

**Service not starting:**
```bash
journalctl -u angel-teamserver -n 50
cat /etc/angel.env  # Verify env vars
```

**TLS certificate error:**
```bash
certbot certificates
certbot renew
```

**Port already in use:**
```bash
ss -tlnp | grep <port>
kill $(lsof -ti:<port>)
```

**Implant not connecting:**
```bash
# Check firewall
ufw status

# Check DNS
dig <domain>

# Check TLS
curl -vk https://<domain>:8443
```