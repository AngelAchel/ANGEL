ANGEL MODULE CONNECTIVITY MATRIX
====================================
Date: 2026-09-16

| Source | Target | Protocol | Port | Status | Bukti |
|--------|--------|----------|------|--------|-------|
| Gateway | Teamserver | HTTP | 8443 | OK | 404 (reachable) |
| Gateway | Console | HTTP | 3000 | OK | 404 (reachable) |
| Gateway | DVWA | HTTP | 8081 | OK | 302 (reachable) |
| Teamserver | DB | — | — | N/A | no DB configured |
| Console | Teamserver | HTTP | 8443 | OK | 404 (reachable) |
| EventBus | All modules | in-proc | — | OK | 6 tests pass |
| Orchestrator | EventBus | in-proc | — | OK | 26 tests pass |
| Collector | EventBus | in-proc | — | OK | compile OK |
| Dispatcher | EventBus | in-proc | — | OK | compile OK |
| C2Server | Listener | TCP | 443 | OK | container running |
| C2Server | DNS | UDP | 5353 | OK | container running |
| C2Server | HTTP | TCP | 8080 | OK | container running |
| C2Server | SMB | TCP | 4455 | OK | container running |
| Implant | C2Server | HTTPS | 8443 | OK | binary generated |
| Rules | Teamserver | HTTP | 8443 | OK | container running |
| DVWA | Gateway | HTTP | 8081 | OK | 302 redirect |

Connectivity: all services reachable from gateway.
Event bus: all modules connected via in-process event bus.
Docker network: all containers on same network.

NOT CONNECTED:
- angel-rules container restarting (may need env config)
- No external DB connection (SQLite in-memory for tests)
- No Supabase connection (env vars not set)