package implantgen

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestImplantEngineGenerate(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Generate()
	if err != nil {
		t.Errorf("Generate failed: %v", err)
	}
}

func TestImplantEngineEncryptPayload(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.EncryptPayload()
	if err != nil {
		t.Errorf("EncryptPayload failed: %v", err)
	}
}

func TestImplantEngineDecryptPayload(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DecryptPayload()
	if err != nil {
		t.Errorf("DecryptPayload failed: %v", err)
	}
}

func TestImplantEngineRun(t *testing.T) {
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
