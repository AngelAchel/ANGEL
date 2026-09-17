package integration

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/angel-platform/angel/gateway"
)

// getFreePort returns an available TCP port.
func getFreePort() int {
	l, _ := net.Listen("tcp", "localhost:0")
	if l == nil {
		return 39999 // fallback
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()
	return port
}

// StartGateway starts the gateway in a goroutine and returns a cleanup
// function that properly shuts down the server.
func StartGateway(t testing.TB, gw *gateway.Gateway) func() {
	t.Helper()

	port := gw.Config().Port
	if port == 0 || port == 3000 {
		port = getFreePort()
		gw.Config().Port = port
		gw.SetPort(port)
	}
	startErr := make(chan error, 1)

	go func() {
		startErr <- gw.Start()
	}()

	addr := fmt.Sprintf("http://localhost:%d/api/v1/health", port)
	ready := make(chan struct{})
	go func() {
		defer close(ready)
		for i := 0; i < 50; i++ {
			resp, err := http.Get(addr)
			if err == nil {
				io.Copy(io.Discard, resp.Body)
				resp.Body.Close()
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-ready:
		// Server is ready
	case err := <-startErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("gateway failed to start: %v", err)
		}
	}

	return func() {
		stopErr := gw.Stop()
		if stopErr != nil && !errors.Is(stopErr, http.ErrServerClosed) {
			t.Logf("gateway stop warning: %v", stopErr)
		}
		select {
		case err := <-startErr:
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				t.Logf("gateway start returned: %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Log("gateway did not stop within timeout")
		}
	}
}
