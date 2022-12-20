package types

func ToPtr[T any](in T) *T {
	return &in
}
