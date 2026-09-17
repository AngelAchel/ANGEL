package mdns

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineMdnsSpoof(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.MdnsSpoof()
	if err != nil {
		t.Errorf("MdnsSpoof failed: %v", err)
	}
}

func TestEnginebuildMDNSResponse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.buildMDNSResponse()
	if err != nil {
		t.Errorf("buildMDNSResponse failed: %v", err)
	}
}

func TestEngineLLMNRPoison(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.LLMNRPoison()
	if err != nil {
		t.Errorf("LLMNRPoison failed: %v", err)
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
