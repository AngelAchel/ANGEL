package eventbus

// EventType represents the type of event
type EventType string

const (
	EventTypeEvent   EventType = "event"
	EventTypeCommand EventType = "command"
	EventTypeResult  EventType = "result"
	EventTypeSync    EventType = "sync"
)

// Domain represents the event domain
type Domain string

const (
	DomainC2            Domain = "c2"
	DomainOrchestrator  Domain = "orchestrator"
	DomainExploit       Domain = "exploit"
	DomainInfra         Domain = "infra"
	DomainOps           Domain = "ops"
	DomainReport        Domain = "report"
	DomainGateway       Domain = "gateway"
	DomainFrontend      Domain = "frontend"
)

// Topic represents an event topic
type Topic struct {
	Domain  Domain
	Module  string
	Action  string
	Version string
}

func (t Topic) String() string {
	return string(t.Domain) + "." + t.Module + "." + t.Action + "." + t.Version
}
