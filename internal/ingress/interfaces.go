package ingress

// Validator validates inbound ingress events
type Validator interface {
    Validate(event *IngressEvent) error
}

// Normalizer transforms validated ingress events into canonical form
type Normalizer interface {
    Normalize(event *IngressEvent) (*IngressEvent, error)
}

// Dispatcher hands normalized events to downstream systems
type Dispatcher interface {
    Dispatch(event *IngressEvent) error
}

// Processor defines the full ingress lifecycle
type Processor interface {
    Process(event *IngressEvent) error
}
