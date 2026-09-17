package c2server

import (
	"time"
)

type Node struct{}

func NewNode() *Node {
	return &Node{}
}

func (e *Node) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "node:done")
	return results, nil
}

func (e *Node) Name() string { return "Node" }
func (e *Node) Timestamp() time.Time { return time.Now() }
