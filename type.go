package gr_variable

import "time"

type WriteOnlyGrChannel[T any] interface {
	WriteValue(value T) bool
	MustWriteValue(value T)
	WriteAllValue(values []T) (int, bool)
	MustWriteAllValue(values []T)
	StopWriting()
}

type ReadOnlyGrChannel[T any] interface {
	Receive() <-chan T
	ReadValue() (T, bool)
	ReadAllAvailableValues() ([]T, bool)
	ReadFirstNValues(count int) ([]T, bool)
	ReadAllValues() []T
	ReadAllValuesWithTimeout(timeout time.Duration) ([]T, bool)
}

type GrChannel[T any] interface {
	ReadOnlyGrChannel[T]
	WriteOnlyGrChannel[T]
}
