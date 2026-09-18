package graphql

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineDeepNestedQuery(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DeepNestedQuery()
	if err != nil {
		t.Errorf("DeepNestedQuery failed: %v", err)
	}
}

func TestEngineBatchQueryAbuse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.BatchQueryAbuse()
	if err != nil {
		t.Errorf("BatchQueryAbuse failed: %v", err)
	}
}

func TestEngineSchemaLeak(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SchemaLeak()
	if err != nil {
		t.Errorf("SchemaLeak failed: %v", err)
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
