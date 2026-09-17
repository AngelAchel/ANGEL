package saml

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineSAMLXMLInject(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SAMLXMLInject()
	if err != nil {
		t.Errorf("SAMLXMLInject failed: %v", err)
	}
}

func TestEnginebuildInjectedSAMLResponse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.buildInjectedSAMLResponse()
	if err != nil {
		t.Errorf("buildInjectedSAMLResponse failed: %v", err)
	}
}

func TestEnginebuildAttributeStatement(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.buildAttributeStatement()
	if err != nil {
		t.Errorf("buildAttributeStatement failed: %v", err)
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
