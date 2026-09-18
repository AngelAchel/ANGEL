package api

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineOAuthRedirectAttack(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.OAuthRedirectAttack()
	if err != nil {
		t.Errorf("OAuthRedirectAttack failed: %v", err)
	}
}

func TestEngineJWTAlgorithmBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.JWTAlgorithmBypass()
	if err != nil {
		t.Errorf("JWTAlgorithmBypass failed: %v", err)
	}
}

func TestEngineRateLimitBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.RateLimitBypass()
	if err != nil {
		t.Errorf("RateLimitBypass failed: %v", err)
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
