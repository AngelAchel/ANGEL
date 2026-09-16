ANGEL REGRESSION RESULTS
====================================
Date: 2026-09-16

========================================
FULL REGRESSION TEST
========================================
Command: go test -count=1 ./...
Result: 92 packages ALL PASS
Exit Status: 0

Passing packages (92):
  c2/generate, c2/implant, c2/profiles
  gateway, gateway/auth, gateway/config, gateway/middleware
  modules/layer01-05/c2implant, c2listener, c2server, channel_rotation, dbpost, environment_detection, listener, nosql, resilience, server, sqli
  modules/layer06-10/evasion, kerberos, lateral, persistence, rootkit
  modules/layer11-15/brain, collector, credential, destruction, orchestrator
  modules/layer16-21/cleanup, evidence, exploit, infra, osint, report
  modules/layer22-25/authbypass, destructionchain, implantgen, netevasion
  modules/layer26-40/ai, api, cloud, container, ir, malware, mobile, physical, purpleteam, socialengineering, supplychain, threatintel, web3, wireless, zerotrust
  modules/layer41-60/cachesmuggle, certforgery, compliance, csrf, dnssec, iot, ipv6, ldap, mdns, methodology, multicloud, opsec, redirect, saml, scada, tls13, upload, webmisc
  modules/layer61-70/arpdhcp, bizlogic, crypto, deser, graphql, grpc, memcorrupt, passwordreset, racecond, vlan
  orchestrator, orchestrator/fireteam, orchestrator/intent, orchestrator/langgraph, orchestrator/mcp
  pkg/crypto, pkg/eventbus, pkg/logger, pkg/types
  tests/e2e, tests/integration

FAILED: 0

========================================
REGRESSION BY CATEGORY
========================================
| Category | Packages | Result |
|----------|----------|--------|
| Build | 83 | ALL PASS |
| Unit Test | 92 | ALL PASS |
| Integration Test | 2 (e2e, integration) | ALL PASS |
| E2E Test | 4 | ALL PASS |
| Lab Script | 25 | ALL PASS |
| Docker | 4 containers | 3 running, 1 restarting |
| Health Check | 4 services | all responding |
| Race Test | 10 | ALL PASS |
| Auth Test | 6 | ALL PASS |

========================================
NO REGRESSION DETECTED
========================================
All previously fixed modules still pass.
No new failures introduced.