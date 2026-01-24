package ingress

import "time"

// Inbound represents raw input entering the ingress system.
type Inbound struct {
	ID        string            // unique identifier for the inbound item
	Source    string            // origin of the input (api, stream, file, etc.)
	Received  time.Time         // time ingress accepted the input
	Payload   []byte            // raw, untrusted data
	Headers   map[string]string // optional transport metadata
}

// Normalized represents validated and canonicalized input.
type Normalized struct {
	ID         string            // propagated from Inbound
	Source     string            // propagated from Inbound
	Received   time.Time         // original receipt time
	Normalized time.Time         // time normalization completed
	Data       any               // canonical internal representation
	Metadata   map[string]string // enrichment for downstream systems
}
