package chan_variable

import (
	"time"
)

type ChanVar[T any] chan T

func (ch ChanVar[T]) WriteValue(value T) {
	ch <- value
}

func (ch ChanVar[T]) WriteAllValue(values []T) {
	for _, item := range values {
		ch.WriteValue(item)
	}
}

func (ch ChanVar[T]) ReadValue() (T, bool) {
	item, ok := <-ch
	return item, ok
}

func (ch ChanVar[T]) ReadAllAvailableValues() ([]T, bool) {
	if cap(ch) == 0 {
		return ch.ReadFirstNValues(1)
	}

	return ch.ReadFirstNValues(len(ch))
}

func (ch ChanVar[T]) ReadFirstNValues(count int) ([]T, bool) {
	result := make([]T, count)

	for i := 0; i < count; i++ {
		if item, ok := ch.ReadValue(); !ok {
			return result[:i], false
		} else {
			result[i] = item
		}
	}

	return result, true
}

func (ch ChanVar[T]) ReadAllValues() []T {
	collection := make([]T, 0)

	for item := range ch {
		collection = append(collection, item)
	}

	return collection
}

func (ch ChanVar[T]) ReadAllValuesWithTimeout(timeout time.Duration) ([]T, bool) {
	expire := time.NewTimer(timeout)
	defer expire.Stop()

	result := make([]T, 0)
	for {
		select {
		case <-expire.C:
			return result, true
		case item, ok := <-ch:
			if !ok {
				return result, false
			}

			result = append(result, item)
		}
	}
}

func (ch ChanVar[T]) StopWriting() {
	close(ch)
}

func NewCharVar[T any]() ChanVar[T] {
	return make(ChanVar[T])
}

func NewCharVarWithLength[T any](length int) ChanVar[T] {
	return make(ChanVar[T], length)
}
