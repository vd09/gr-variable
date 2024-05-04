package gr_variable

import (
	"testing"
)

func TestReadAllAvailableValues(t *testing.T) {
	t.Run("TestNewGoRoutineVariable", testNewGrChannel)
}

func testNewGrChannel(t *testing.T) {
	t.Run("TestReadAllValues", testReadAllValues)
	t.Run("TestReadFromEmptyChannel", testReadFromEmptyChannel)
}

func testReadAllValuesWrites(ch WriteOnlyGrChannel[int], values []int) {
	go func() {
		// Write some values to the channel
		ch.WriteAllValue(values)
	}()
}

func testReadAllValuesReadAndValidate(t *testing.T, ch ReadOnlyGrChannel[int], values []int) {
	totalElements := 0
	for totalElements < len(values) {
		// Read all available values
		readValues, ok := ch.ReadAllAvailableValues()

		// Check if all values were read successfully
		if !ok {
			t.Error("Failed to read all available values")
		}

		// Check if all values were read successfully
		if len(readValues) != 1 {
			t.Error("unbuffred channel should only read 1 value at the time")
		}

		// Compare the read values with the original values
		for i, v := range readValues {
			if readValues[i] != values[totalElements] {
				t.Errorf("Value mismatch at index %d, expected: %d, got: %d", i, v, readValues[i])
			}
			totalElements++
		}
	}
}

func testReadAllValues(t *testing.T) {
	// Create a new ChanVar instance
	ch := NewGrChannel[int]()

	values := []int{1, 2, 3}
	testReadAllValuesWrites(ch, values)
	testReadAllValuesReadAndValidate(t, ch, values)
}

func testReadFromEmptyChannel(t *testing.T) {
	// Test with an empty channel
	emptyCh := NewGrChannel[int]()
	emptyCh.StopWriting() // Close the empty channel
	_, emptyOK := emptyCh.ReadAllAvailableValues()

	// Check if reading from an empty channel returns false
	if emptyOK {
		t.Error("Expected ReadAllAvailableValues to return false for an empty channel")
	}
}
