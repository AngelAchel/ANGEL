package c2server

import (
	"time"
)

type Peer struct{}

func NewPeer() *Peer {
	return &Peer{}
}

func (e *Peer) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "peer:done")
	return results, nil
}

func (e *Peer) Name() string         { return "Peer" }
func (e *Peer) Timestamp() time.Time { return time.Now() }
