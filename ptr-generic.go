package ptr

// Ptr returns a new pointer to the given value.
func Ptr[T any](v T) *T {
	return &v
}

// Value returns a value behind the pointer or the zero value when the pointer is nil.
func Value[T any](p *T) T {
	if p != nil {
		return *p
	}

	var zeroValue T
	return zeroValue
}

// ValueOrDefault returns a value behind the pointer or the default value when the pointer is nil.
func ValueOrDefault[T any](p *T, defaultValue T) T {
	if p != nil {
		return *p
	}

	return defaultValue
}
