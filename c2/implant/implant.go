package implant

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"os"
	"runtime"
	"time"
)

type Implant struct {
	config    *ImplantConfig
	running   bool
	sessionID string
}

type ImplantConfig struct {
	ServerURL  string
	Sleep      int
	Jitter     float64
	Profile    string
	Encryption string
}

func NewImplant(cfg *ImplantConfig) *Implant {
	if cfg == nil {
		cfg = &ImplantConfig{
			ServerURL:  "https://teamserver.angel.local",
			Sleep:      30,
			Jitter:     0.25,
			Profile:    "default",
			Encryption: "aes-256-gcm",
		}
	}
	return &Implant{config: cfg}
}

func (i *Implant) Run() {
	i.running = true
	i.sessionID = i.generateSessionID()

	for i.running {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		jitterMs := float64(n.Int64()) / 100.0 * i.config.Jitter * float64(time.Second)
		time.Sleep(time.Duration(i.config.Sleep)*time.Second + time.Duration(jitterMs))
		i.checkIn()
	}
}

func (i *Implant) checkIn() {
	data := map[string]string{
		"session_id":  i.sessionID,
		"hostname":    i.getHostname(),
		"internal_ip": i.getInternalIP(),
		"user":        i.getCurrentUser(),
		"os":          runtime.GOOS,
		"arch":        runtime.GOARCH,
	}
	_ = data
}

func (i *Implant) Stop() {
	i.running = false
}

func (i *Implant) generateSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (i *Implant) getHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func (i *Implant) getInternalIP() string {
	return "127.0.0.1"
}

func (i *Implant) getCurrentUser() string {
	return os.Getenv("USER")
}
