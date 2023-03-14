package validation

// A Validator validates the supplied input.
type Validator[T any] interface {
	Validate(in *T) error
}

// A ValidatorFn validates the supplied input.
type ValidatorFn[T any] func(in *T) error

// Validate the supplied input.
func (fn ValidatorFn[T]) Validate(in *T) error {
	return fn(in)
}

// A Chain runs multiple validations.
type Chain[T any] []Validator[T]

// Validate the supplied input.
func (vs Chain[T]) Validate(in *T) error {
	for _, v := range vs {
		if err := v.Validate(in); err != nil {
			return err
		}
	}
	return nil
}
