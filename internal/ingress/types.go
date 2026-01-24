package ingress

// IngressEvent represents a unit of external input entering the system.
type IngressEvent struct {
	ID        string
	Source    string
	Timestamp int64
	Payload   any
	Metadata  map[string]string
}
