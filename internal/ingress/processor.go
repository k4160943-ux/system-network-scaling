package ingress

// Processor coordinates the ingress lifecycle:
// validate -> normalize -> dispatch.
type Processor struct {
	Validator  Validator
	Normalizer Normalizer
	Dispatcher Dispatcher
}

// Process executes the ingress lifecycle for a single event.
func (p *Processor) Process(event *IngressEvent) error {
	if err := p.Validator.Validate(event); err != nil {
		return err
	}

	normalized, err := p.Normalizer.Normalize(event)
	if err != nil {
		return err
	}

	if err := p.Dispatcher.Dispatch(normalized); err != nil {
		return err
	}

	return nil
}
