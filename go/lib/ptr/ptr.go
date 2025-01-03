package ptr

func Ptr[T any](in T) *T {
	return &in
}

func Value[T any](in *T) T {
	if in == nil {
		var fb T
		return fb
	}
	return *in
}

func ValueOrDefault[T any](in *T, fallback T) T {
	if in == nil {
		return fallback
	}
	return *in
}
