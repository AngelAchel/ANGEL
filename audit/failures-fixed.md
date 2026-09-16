ANGEL FAILURES FIXED
====================================
Date: 2026-09-16

========================================
ERRCHECK FIXES (30+ FILES)
========================================
| File | Issue | Fix |
|------|-------|-----|
| listener/dns_listener.go | defer ln.Close() | defer func() { _ = ln.Close() }() |
| listener/listener.go | conn.Close() | _ = conn.Close() |
| listener/tcp_listener.go | defer ln.Close() | defer func() { _ = ln.Close() }() |
| listener/http_listener.go | fmt.Fprintf | _, _ = fmt.Fprintf(...) |
| listener/udp_listener.go | defer ln.Close() | defer func() { _ = ln.Close() }() |
| c2server/listener.go | eb.Publish, w.Write, Shutdown, SetReadDeadline, WriteTo, conn.Close | _, _ = ... |
| c2server/teamserver.go | eb.Publish | _, _ = ... |
| decoy/decoy.go | fmt.Fprintf, r.ParseForm | _, _ = ... |
| decoy/honeypot.go | r.ParseForm | _ = r.ParseForm() |
| self_destruct.go | os.Remove | _ = os.Remove(...) |
| firmware.go | cmd.CombinedOutput | _, _ = cmd.CombinedOutput() |
| smm.go | cmd.CombinedOutput, os.RemoveAll | _, _ = ... |
| cleanup_test.go | os.WriteFile, os.MkdirAll | _ = ... |
| iot/engine_test.go | CredentialDump, HardcodedKey, BackdoorDetect | _, _ = ... |
| ldap/engine.go | fmt.Sscanf | _, _ = fmt.Sscanf(...) |
| ios_client.go | defer resp.Body.Close() | defer func() { _ = ... }() |
| manager.go | defer resp.Body.Close() | defer func() { _ = ... }() |
| c2_channel_test.go | ch.Reconnect, ch.Connect | _ = ... |
| encryption.go | e.SetKey | _ = e.SetKey(...) |
| mssql.go | m.executeQuery | _, _ = m.executeQuery(...) |
| antianalysis.go | fmt.Sscanf | _, _ = fmt.Sscanf(...) |
| implant_linux.go | defer resp.Body.Close(), os.Remove | defer func()(), _ = os.Remove() |
| detectors.go | defer resp.Body.Close() | defer func() { _ = ... }() |
| sqli/engine_test.go | fmt.Fprintf | _, _ = fmt.Fprintf(...) |
| dns.go | defer ln.Close() | defer func() { _ = ln.Close() }() |
| websocket.go | defer conn.Close(), WriteMessage | defer func()(), _, _ = ... |
| c2listener/listener.go | fmt.Fprintf | _, _ = fmt.Fprintf(...) |
| android.go | os.Remove | _ = os.Remove(...) |
| darwin.go | cmd.CombinedOutput | _, _ = cmd.CombinedOutput() |
| windows.go | startCmd/stopCmd.CombinedOutput | _, _ = ... |
| collector_test.go | kl.Stop, cm.Stop, engine.Collect | _ = ..., _, _ = ... |
| collector/engine.go | keylog.Start, clip.Start | _ = ... |
| ransomware.go | r.encryptFile | _ = r.encryptFile(...) |
| dispatcher.go | d.bus.Publish | _, _ = d.bus.Publish(...) |
| orchestrator/engine.go | o.state.SaveState | _ = o.state.SaveState(...) |
| orchestrator_test.go | sm.SaveState | _ = sm.SaveState(...) |
| authbypass/engine.go | resp2.Body.Close, probeResp.Body.Close | _ = ... |
| supabase/client.go | defer resp.Body.Close() | defer func() { _ = ... }() |
| sqli/detectors.go | defer resp.Body.Close() | defer func() { _ = ... }() |
| environment_test.go | e.Patch, DisableETW, EnableETW, RestoreHooks, os.WriteFile, os.Chtimes, ts.Touch, ts.RandomizeTime, ts.SetTime | _ = ... |

========================================
BROKEN SYNTAX FIXED (10+ FILES)
========================================
| File | Issue | Fix |
|------|-------|-----|
| firmware.go | "if output, err := _, _ = cmd.CombinedOutput()" | Split into two lines |
| smm.go | Same pattern | Split into two lines |
| cleanup_test.go | "_, _ = os.WriteFile" | "_ = os.WriteFile" |
| iot/engine_test.go | "result, err := _ = eng.HardcodedKey" | Proper declaration |
| ldap/engine.go | "_, _ = _, _ = fmt.Sscanf" | Proper assignment |
| listener.go | "l._ = conn.Close()" | "_ = l.conn.Close()" |
| teamserver.go | Broken Publish | Proper assignment |
| decoy.go | Broken Fprintf | Proper assignment |
| http_listener.go | Broken Fprintf | Proper assignment |
| c2_channel_test.go | "err := _, _ = ch.Connect" | Proper declaration |
| encryption.go | "_, _ = e.SetKey" | "_ = e.SetKey" |
| mssql.go | "_, err := _, _ = m.executeQuery" | Proper assignment |
| antianalysis.go | Broken Sscanf | Proper assignment |
| environment_test.go | "if err := _ = e.DisableETW" | Proper declaration |
| implant_linux.go | "defer _, _ = resp.Body.Close()" | defer func() |
| udp_listener.go | "defer _ = ln.Close()" | defer func() |
| detectors.go | "defer _, _ = resp.Body.Close()" | defer func() |
| sqli/engine_test.go | Broken Fprintf | Proper assignment |
| dns.go | "defer _ = ln.Close()" | defer func() |
| websocket.go | "return _, _ = conn.WriteMessage" | err := ...; return err |
| c2listener/listener.go | Broken Fprintf | Proper assignment |
| collector_test.go | "if err := _ = kl.Stop()" | Proper declaration |
| orchestrator_test.go | "err := _, _ = sm.SaveState" | Proper declaration |
| darwin.go | Broken CombinedOutput | Split into two lines |
| windows.go | Broken CombinedOutput | Split into two lines |

========================================
TEST FIXES
========================================
| File | Issue | Fix |
|------|-------|-----|
| collector_test.go line 94 | if err := _ = kl.Stop() | if err := kl.Stop() |
| orchestrator_test.go line 255 | err := _, _ = sm.SaveState | err := sm.SaveState |
| orchestrator_test.go line 279-326 | _, _ = sm.SaveState | _ = sm.SaveState |
| collector_test.go line 358 | _, err := _, _ = engine.Collect | _, err := engine.Collect |

========================================
UNUSED CODE DOCUMENTED (~20 FUNCTIONS)
========================================
- persistCron, persistSystemd, persistSSHKeys, persistPAM, cleanupLogs, injectProcess, ptraceInject, procMemWrite, forkBomb (implant_linux.go)
- generateRandomString, calculateVariance, calculateStdDev (traffic_gen.go)
- fileExists (edr_detect.go)
- generateMalleableURL (profile_apply.go)
- formatResult (result_handler.go)
- running, formatTaskStatus (task_queue.go)
- formatBeaconStatus (beacon_core.go)
- totalWidth (payloads.go)
- tamper (types.go)
- mu (adcs.go, adrecon.go)
- generateRandomBytes (ticket.go)
- getAndroidEnvVar, getLinuxEnvVar, getLinuxUserHome (persistence)
- decryptDPAPI (browser.go)
- wireGuardTemplate, wgData, openVPNTemplate, digCommand, extractMetaTags (osint, vpn)
- randomDelay (bruteforce.go)
- timeNow (session.go)
- generateRandomPerturbation, hashPrompt (ai/engine.go)
- generateHash (container/engine.go)
- enumerateKeychainItems, analyzeSSLChain (mobile/engine.go)
- analyzeBytecode (web3/engine.go)
- signalToDistance (wireless/engine.go)
- buildCertTemplate (certforgery/engine.go)
- calculateKeySize, hammingDistance, detectCipherMode, analyzeIV (crypto/engine.go)
- formatTags (vlan/engine.go)

========================================
STATICCHECK WARNINGS (~60)
========================================
- Deprecated crypto APIs (ScalarMult, Marshal, Unmarshal)
- Empty branch
- Unnecessary fmt.Sprintf
- Deprecated math/rand.Read
- Tagged switch suggestions
- Capitalized error strings
- Unused append result
- Should use append pattern