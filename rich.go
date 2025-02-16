package enum

// RichEnum provides a set of utility methods to simplify working with enums.
//
// It includes various helper functions for operations like serialization,
// deserialization, string conversion, and validation, making it easier to
// manage and manipulate enum values across your codebase.
type RichEnum[T any] int

func (e RichEnum[T]) IsValid() bool {
	return IsValid(e)
}

func (e RichEnum[T]) MarshalJSON() ([]byte, error) {
	return MarshalJSON(e)
}

func (e *RichEnum[T]) UnmarshalJSON(data []byte) error {
	return UnmarshalJSON(data, e)
}

func (e RichEnum[T]) Int() int {
	return int(e)
}

func (e RichEnum[T]) String() string {
	return StringOf(e)
}
