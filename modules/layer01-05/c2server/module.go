package c2server

import (
	"fmt"
	"sync"
	"time"

	"github.com/angel-platform/angel/pkg/eventbus"
	"github.com/angel-platform/angel/pkg/logger"
	"github.com/angel-platform/angel/pkg/types"
)

type Module struct {
	mu       sync.RWMutex
	eb       *eventbus.EventBus
	log      *logger.Logger
	handlers map[string]TaskHandler
}

type TaskHandler func(task *types.Task) ([]byte, error)

func NewModule() *Module {
	eb := eventbus.New("module-key")
	log := logger.New("c2-module", logger.LevelInfo)
	return &Module{
		eb:       eb,
		log:      log,
		handlers: make(map[string]TaskHandler),
	}
}

func (e *Module) RegisterTaskHandler(name string, handler TaskHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.handlers[name] = handler
}

func (e *Module) Run() ([]string, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	results := make([]string, 0, len(e.handlers)+1)
	results = append(results, fmt.Sprintf("module:active:tasks=%d", len(e.handlers)))
	for name := range e.handlers {
		results = append(results, fmt.Sprintf("task:%s:registered", name))
	}
	return results, nil
}

func (e *Module) Name() string         { return "C2Module" }
func (e *Module) Timestamp() time.Time { return time.Now() }

func (e *Module) ExecuteTask(task *types.Task) ([]byte, error) {
	e.mu.RLock()
	handler := e.handlers[string(task.Type)]
	e.mu.RUnlock()
	if handler == nil {
		return nil, fmt.Errorf("no handler for task type: %s", task.Type)
	}
	return handler(task)
}
