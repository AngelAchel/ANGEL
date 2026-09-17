package bizlogic

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEnginePriceManipulation(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PriceManipulation()
	if err != nil {
		t.Errorf("PriceManipulation failed: %v", err)
	}
}

func TestEngineQuantityNeg(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.QuantityNeg()
	if err != nil {
		t.Errorf("QuantityNeg failed: %v", err)
	}
}

func TestEngineCouponAbuse(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.CouponAbuse()
	if err != nil {
		t.Errorf("CouponAbuse failed: %v", err)
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
