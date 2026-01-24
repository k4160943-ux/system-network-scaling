package ingress

// ValidatorFunc allows validators to be implemented as functions.
type ValidatorFunc func(event *IngressEvent) error

// Validate executes the validator function.
func (vf ValidatorFunc) Validate(event *IngressEvent) error {
	return vf(event)
}
