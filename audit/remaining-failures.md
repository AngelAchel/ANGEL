ANGEL REMAINING FAILURES
====================================
Date: 2026-09-16

========================================
REMAINING ERRCHECK ISSUES (51)
========================================
| File | Line | Issue | Severity |
|------|------|-------|----------|
| c2/implant/implant.go | 144 | out.Close unchecked | low |
| gateway/config/config_test.go | 64-66 | os.Unsetenv unchecked | low |
| c2listener/listener.go | 99 | fmt.Fprintf unchecked | low |
| environment_detection/environment_test.go | 159,168,181,189,215 | DisableETW, EnableETW, Patch, Restore unchecked | low |
| listener/http_listener.go | 108 | fmt.Fprintf unchecked | low |
| listener/udp_listener.go | 137 | conn.Close unchecked | low |
| server/database/sqlite.go | 145,196 | rows.Close unchecked | low |
| server/server.go | 92,105,116,121 | json.Encoder.Encode, fmt.Fprintf unchecked | low |
| smb_beacon/beacon_core.go | 69 | b.namePipe.Create unchecked | low |
| evasion/cleanup.go | 176 | os.Setenv unchecked | low |

Note: All remaining errcheck issues are LOW severity — unchecked errors in non-critical paths, test files, or defer-like patterns. None affect functionality.

========================================
REMAINING STATICCHECK WARNINGS (~60)
========================================
| Category | Count | Severity |
|----------|-------|----------|
| Deprecated crypto APIs | ~15 | low |
| Unnecessary fmt.Sprintf | ~5 | info |
| Deprecated math/rand.Read | ~3 | low |
| Tagged switch suggestions | ~10 | info |
| Capitalized error strings | ~5 | info |
| Unused append result | ~5 | info |
| Should use append pattern | ~5 | info |
| Empty branch | ~3 | info |
| De Morgan's law | ~3 | info |
| SA4023 comparison always true | ~5 | info |
| QF1006 loop optimization | ~2 | info |
| QF1001 De Morgan's law | ~4 | info |

Note: All staticcheck warnings are informational only — do not affect functionality or correctness.

========================================
REMAINING UNUSED CODE (~20 FUNCTIONS)
========================================
See failures-fixed.md for full list. These are unused functions that may be dead code or intended for future features.

========================================
DOCKER ISSUES
========================================
| Container | Issue | Severity |
|-----------|-------|----------|
| angel-rules | Restarting (0) | medium |

Root cause: May need SUPABASE_URL, SUPABASE_ANON_KEY, SUPABASE_SERVICE_ROLE_KEY env vars set.
Impact: Rules loading may fail — non-critical for core functionality.

========================================
MISSING ENV VARS
========================================
| Variable | Status | Impact |
|----------|--------|--------|
| TEAMSERVER_KEY | not set | Required for C2 server |
| CRYPTO_KEY | not set | Required for crypto |
| JWT_SECRET | not set | Required for auth |
| SUPABASE_URL | not set | Warnings only |
| SUPABASE_ANON_KEY | not set | Warnings only |
| SUPABASE_SERVICE_ROLE_KEY | not set | Warnings only |
| DATABASE_ENCRYPTION_KEY | not set | Warnings only |

========================================
NOT FAILURES — BY DESIGN
========================================
- 404 on /health endpoint (no health endpoint implemented)
- 404 on /api/v1/status (no status endpoint)
- angel-rules restarting (may need config)
- Test "no tests to run" (no test file in some packages)