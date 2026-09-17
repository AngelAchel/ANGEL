package container

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineDockerEscape(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DockerEscape()
	if err != nil {
		t.Errorf("DockerEscape failed: %v", err)
	}
}

func TestEngineKubernetesAPIAccess(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.KubernetesAPIAccess()
	if err != nil {
		t.Errorf("KubernetesAPIAccess failed: %v", err)
	}
}

func TestEngineExtractSecrets(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ExtractSecrets()
	if err != nil {
		t.Errorf("ExtractSecrets failed: %v", err)
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
