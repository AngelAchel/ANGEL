package iot

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineFirmwareExtract(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FirmwareExtract()
	if err != nil {
		t.Errorf("FirmwareExtract failed: %v", err)
	}
}

func TestEngineanalyzeFirmware(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.analyzeFirmware()
	if err != nil {
		t.Errorf("analyzeFirmware failed: %v", err)
	}
}

func TestEngineformatFirmwareInfo(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatFirmwareInfo()
	if err != nil {
		t.Errorf("formatFirmwareInfo failed: %v", err)
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
