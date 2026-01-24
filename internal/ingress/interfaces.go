package ingress

// Validator defines the contract for ingress validation.
type Validator interface {
	Validate(event *IngressEvent) error
}

// Normalizer defines the contract for transforming ingress events
// into canonical internal representations.
type Normalizer interface {
	Normalize(event *IngressEvent) (*IngressEvent, error)
}

// Dispatcher defines the contract for handing ingress events
// to downstream systems.
type Dispatcher interface {
	Dispatch(event *IngressEvent) error
}
