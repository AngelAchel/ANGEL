package supplychain

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineNPMTyposquat(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.NPMTyposquat()
	if err != nil {
		t.Errorf("NPMTyposquat failed: %v", err)
	}
}

func TestEngineGitHubActionsInject(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.GitHubActionsInject()
	if err != nil {
		t.Errorf("GitHubActionsInject failed: %v", err)
	}
}

func TestEngineDockerfileInject(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DockerfileInject()
	if err != nil {
		t.Errorf("DockerfileInject failed: %v", err)
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
