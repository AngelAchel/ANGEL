package purpleteam

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineAlertValidation(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.AlertValidation()
	if err != nil {
		t.Errorf("AlertValidation failed: %v", err)
	}
}

func TestEngineDetectionRuleTest(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DetectionRuleTest()
	if err != nil {
		t.Errorf("DetectionRuleTest failed: %v", err)
	}
}

func TestEngineLogCoverageTest(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.LogCoverageTest()
	if err != nil {
		t.Errorf("LogCoverageTest failed: %v", err)
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
