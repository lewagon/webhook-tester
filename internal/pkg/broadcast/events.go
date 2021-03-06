package broadcast

// Event is something that happens (sorry for the tautology).
type Event interface {
	// Name returns event name.
	Name() string

	// Data returns event payload.
	Data() interface{}
}

type event struct {
	name string
	data interface{}
}

// Name returns event name.
func (e *event) Name() string { return e.name }

// Data returns event payload.
func (e *event) Data() interface{} { return e.data }

// NewRequestRegisteredEvent creates an event, that means "new request with passed ID was registered".
func NewRequestRegisteredEvent(requestID string) Event {
	return &event{name: "request-registered", data: requestID}
}

// NewRequestDeletedEvent creates an event, that means "request with passed ID was deleted".
func NewRequestDeletedEvent(requestID string) Event {
	return &event{name: "request-deleted", data: requestID}
}

// NewAllRequestsDeletedEvent creates an event, that means "all requests was deleted".
func NewAllRequestsDeletedEvent() Event {
	return &event{name: "requests-deleted", data: "*"}
}
