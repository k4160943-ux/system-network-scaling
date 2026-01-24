package ingress

// DispatcherFunc allows dispatchers to be implemented as functions.
type DispatcherFunc func(event *IngressEvent) error

// Dispatch executes the dispatcher function.
func (df DispatcherFunc) Dispatch(event *IngressEvent) error {
	return df(event)
}
