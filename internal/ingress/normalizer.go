package ingress

// NormalizerFunc allows normalizers to be implemented as functions.
type NormalizerFunc func(event *IngressEvent) (*IngressEvent, error)

// Normalize executes the normalizer function.
func (nf NormalizerFunc) Normalize(event *IngressEvent) (*IngressEvent, error) {
	return nf(event)
}
