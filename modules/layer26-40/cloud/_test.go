package cloud

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineAWSIAMPrivesc(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.AWSIAMPrivesc()
	if err != nil {
		t.Errorf("AWSIAMPrivesc failed: %v", err)
	}
}

func TestEngineAzureADAttack(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.AzureADAttack()
	if err != nil {
		t.Errorf("AzureADAttack failed: %v", err)
	}
}

func TestEngineGCPServiceAccountAbuse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.GCPServiceAccountAbuse()
	if err != nil {
		t.Errorf("GCPServiceAccountAbuse failed: %v", err)
	}
}

func TestEngineRun(t *testing.T) {
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
