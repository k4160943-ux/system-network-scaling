package ingress

import (
	"errors"
	"testing"
)

func TestProcessor_Process_Success(t *testing.T) {
	called := struct {
		validate  bool
		normalize bool
		dispatch  bool
	}{}

	p := &Processor{
		Validator: ValidatorFunc(func(event *IngressEvent) error {
			called.validate = true
			return nil
		}),
		Normalizer: NormalizerFunc(func(event *IngressEvent) (*IngressEvent, error) {
			called.normalize = true
			return event, nil
		}),
		Dispatcher: DispatcherFunc(func(event *IngressEvent) error {
			called.dispatch = true
			return nil
		}),
	}

	err := p.Process(&IngressEvent{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !called.validate || !called.normalize || !called.dispatch {
		t.Fatalf("expected all stages to be called")
	}
}

func TestProcessor_Process_ValidationFailure(t *testing.T) {
	expectedErr := errors.New("validation failed")

	p := &Processor{
		Validator: ValidatorFunc(func(event *IngressEvent) error {
			return expectedErr
		}),
		Normalizer: NormalizerFunc(func(event *IngressEvent) (*IngressEvent, error) { return event, nil }),
		Dispatcher: DispatcherFunc(func(event *IngressEvent) error { return nil }),
	}

	err := p.Process(&IngressEvent{})
	if err != expectedErr {
		t.Fatalf("expected validation error to propagate")
	}
}
