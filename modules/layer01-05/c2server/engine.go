package c2server

import (
	"fmt"
	"sync"
	"time"

	"github.com/angel-platform/angel/pkg/eventbus"
	"github.com/angel-platform/angel/pkg/logger"
	"github.com/angel-platform/angel/pkg/types"
)

type Engine struct {
	mu       sync.RWMutex
	eb       *eventbus.EventBus
	log      *logger.Logger
	agents   map[string]*types.Agent
	handlers map[string]CommandHandler
}

type CommandHandler func(cmd string) ([]byte, error)

func NewEngine() *Engine {
	eb := eventbus.New("engine-key")
	log := logger.New("c2-engine", logger.LevelInfo)
	return &Engine{
		eb:       eb,
		log:      log,
		agents:   make(map[string]*types.Agent),
		handlers: make(map[string]CommandHandler),
	}
}

func (e *Engine) RegisterHandler(name string, handler CommandHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers[name] = handler
}

func (e *Engine) Run() ([]string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	results := make([]string, 0, len(e.agents)+1)
	results = append(results, fmt.Sprintf("engine:active:agents=%d", len(e.agents)))
	for id := range e.agents {
		results = append(results, fmt.Sprintf("agent:%s:connected", id))
	}
	return results, nil
}

func (e *Engine) Name() string         { return "C2Engine" }
func (e *Engine) Timestamp() time.Time { return time.Now() }

func (e *Engine) ExecuteCommand(agentID string, cmd string) ([]byte, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	handler := e.handlers["shell"]
	if handler == nil {
		return nil, fmt.Errorf("no handler registered")
	}
	return handler(cmd)
}

func (e *Engine) GetAgentCount() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return len(e.agents)
}
