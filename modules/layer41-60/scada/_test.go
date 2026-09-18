package scada

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineModbusEnumerate(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ModbusEnumerate()
	if err != nil {
		t.Errorf("ModbusEnumerate failed: %v", err)
	}
}

func TestEngineenumerateModbus(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.enumerateModbus()
	if err != nil {
		t.Errorf("enumerateModbus failed: %v", err)
	}
}

func TestEngineformatModbusDevices(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatModbusDevices()
	if err != nil {
		t.Errorf("formatModbusDevices failed: %v", err)
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
