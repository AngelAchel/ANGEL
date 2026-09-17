package utility

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// UARotateEngine provides user-agent rotation utilities.
type UARotateEngine struct {
	userAgents []string
	index      int
}

// NewUARotateEngine creates a new UARotateEngine.
func NewUARotateEngine() *UARotateEngine {
	return &UARotateEngine{
		userAgents: defaultUserAgents(),
	}
}

// Next returns the next user-agent in rotation.
func (e *UARotateEngine) Next() string {
	if len(e.userAgents) == 0 {
		return "Mozilla/5.0"
	}
	ua := e.userAgents[e.index%len(e.userAgents)]
	e.index++
	return ua
}

// Random returns a random user-agent string.
func (e *UARotateEngine) Random() string {
	if len(e.userAgents) == 0 {
		return "Mozilla/5.0"
	}
	return e.userAgents[rng.Intn(len(e.userAgents))]
}

// Add appends a user-agent to the rotation list.
func (e *UARotateEngine) Add(ua string) {
	e.userAgents = append(e.userAgents, ua)
}

// Set replaces the entire user-agent list.
func (e *UARotateEngine) Set(uas []string) {
	e.userAgents = uas
	e.index = 0
}

// Reset resets the rotation index.
func (e *UARotateEngine) Reset() {
	e.index = 0
}

// Count returns the number of user-agents in rotation.
func (e *UARotateEngine) Count() int {
	return len(e.userAgents)
}

// Browser generates a browser-specific user-agent.
func (e *UARotateEngine) Browser(browser string) string {
	switch strings.ToLower(browser) {
	case "chrome":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.%d.%d Safari/537.36",
			120+rng.Intn(5), rng.Intn(9999), rng.Intn(99))
	case "firefox":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:%d.0) Gecko/20100101 Firefox/%d.0",
			110+rng.Intn(15), 115+rng.Intn(10))
	case "safari":
		return fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_%d_%d) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%d.%d Safari/605.1.15",
			14+rng.Intn(3), rng.Intn(10), 15+rng.Intn(5), rng.Intn(9))
	case "edge":
		return fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.0.%d.%d Safari/537.36 Edg/%d.0.%d",
			120+rng.Intn(5), rng.Intn(9999), rng.Intn(99), 120+rng.Intn(5), rng.Intn(9999))
	default:
		return e.Random()
	}
}

func defaultUserAgents() []string {
	return []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:121.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Android 14; Mobile; rv:121.0) Gecko/121.0 Firefox/121.0",
		"Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
	}
}
