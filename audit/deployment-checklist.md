ANGEL DEPLOYMENT CHECKLIST
====================================
Date: 2026-09-16
Status: PRODUCTION READY WITH LIMITATIONS

========================================
PREREQUISITES
========================================
[ ] Docker installed (24.x+)
[ ] Docker Compose installed (2.x+)
[ ] Git installed
[ ] Go 1.22+ (for local build)
[ ] 10GB disk space
[ ] 4GB RAM minimum
[ ] Network isolated (lab VLAN)

========================================
ENVIRONMENT VARS (REQUIRED)
========================================
[ ] TEAMSERVER_KEY — set via .env file
[ ] CRYPTO_KEY — set via .env file
[ ] JWT_SECRET — set via .env file
[ ] DATABASE_ENCRYPTION_KEY — set via .env file
[ ] SUPABASE_URL — optional (for cloud rules)
[ ] SUPABASE_ANON_KEY — optional (for cloud rules)
[ ] SUPABASE_SERVICE_ROLE_KEY — optional (for cloud rules)

========================================
DOCKER DEPLOYMENT
========================================
[ ] git clone repo
[ ] cp .env.example .env
[ ] edit .env with real values (TEAMSERVER_KEY, CRYPTO_KEY, JWT_SECRET)
[ ] docker-compose build
[ ] docker-compose up -d
[ ] docker-compose ps (verify all containers Up)
[ ] docker-compose logs -f (verify no errors)

========================================
HEALTH CHECK
========================================
[ ] curl http://localhost:8443/ → 404 (teamserver OK)
[ ] curl http://localhost:3000/ → 404 (console OK)
[ ] curl http://localhost:8081/ → 302 (DVWA OK)
[ ] docker ps (all containers Up)
[ ] docker-compose ps (all services healthy)

========================================
VERIFICATION
========================================
[ ] go build ./... (compile OK)
[ ] go test -count=1 ./... (92 PASS)
[ ] bash lab/test_lab_full.sh (25 PASS)
[ ] go test -race -count=1 ./... (no race conditions)
[ ] golangci-lint run --no-config ./... (0 errcheck)

========================================
PRODUCTION HARDENING
========================================
[ ] Change default secrets
[ ] Set real TEAMSERVER_KEY, CRYPTO_KEY, JWT_SECRET
[ ] Enable TLS certificates
[ ] Configure firewall rules
[ ] Set up logging/monitoring
[ ] Configure backup strategy
[ ] Set resource limits (memory/CPU)
[ ] Enable Docker restart policy
[ ] Configure network isolation
[ ] Review unused code (~20 functions)

========================================
DO NOT DEPLOY WITHOUT
========================================
[ ] Explicit production instruction
[ ] Real credentials in .env (not .env.example)
[ ] TLS certificates configured
[ ] Firewall rules set
[ ] Backup strategy in place
[ ] Monitoring/logging configured

========================================
WARNING
========================================
- This is an offensive security platform
- Only use on authorized systems
- Do not attack systems without permission
- Keep credentials secure
- Rotate keys regularly
- Do not commit .env to git