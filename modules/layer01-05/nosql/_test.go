package nosql

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestNoSQLEngineScan(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Scan()
	if err != nil {
		t.Errorf("Scan failed: %v", err)
	}
}

func TestNoSQLEngineExploit(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Exploit()
	if err != nil {
		t.Errorf("Exploit failed: %v", err)
	}
}

func TestNoSQLEnginescanMongoDB(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.scanMongoDB()
	if err != nil {
		t.Errorf("scanMongoDB failed: %v", err)
	}
}

func TestNoSQLEngineRun(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	result, err := gw.Run()
	if err != nil {
		t.Errorf("Run failed: %v", err)
	}
	if result == "" {
		t.Error("expected non-empty result")
	}
}
