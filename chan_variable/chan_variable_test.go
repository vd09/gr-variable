package chan_variable

import (
	"testing"
	"time"
)

func TestReadAllAvailableValues(t *testing.T) {
	t.Run("TestNewGoRoutineVariable", testNewGoRoutineVariable)
	t.Run("TestNewGoRoutineVariableWithLength", testNewGoRoutineVariableWithLength)
}

func testNewGoRoutineVariable(t *testing.T) {
	t.Run("TestReadAllValues", testReadAllValues)
	t.Run("TestReadFromEmptyChannel", testReadFromEmptyChannel)
}

func testReadAllValues(t *testing.T) {
	// Create a new ChanVar instance
	ch := NewCharVar[int]()

	values := []int{1, 2, 3}
	go func() {
		// Write some values to the channel
		ch.WriteAllValue(values)
	}()

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

func testReadFromEmptyChannel(t *testing.T) {
	// Test with an empty channel
	emptyCh := NewCharVar[int]()
	emptyCh.StopWriting() // Close the empty channel
	_, emptyOK := emptyCh.ReadAllAvailableValues()

	// Check if reading from an empty channel returns false
	if emptyOK {
		t.Error("Expected ReadAllAvailableValues to return false for an empty channel")
	}
}

func testNewGoRoutineVariableWithLength(t *testing.T) {
	t.Run("TestReadAllValues", testReadAllValuesWithLength)
	t.Run("TestReadFromEmptyChannel", testReadFromEmptyChannelWithLength)
}

func testReadAllValuesWithLength(t *testing.T) {
	// Create a new ChanVar instance
	ch := NewCharVarWithLength[int](5)

	values := []int{1, 2, 3}
	go func() {
		// Write some values to the channel
		ch.WriteAllValue(values)
	}()

	totalElements := 0
	for totalElements < len(values) {
		// Read all available values
		readValues, ok := ch.ReadAllAvailableValues()

		// Check if all values were read successfully
		if !ok {
			t.Error("Failed to read all available values")
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

func testReadFromEmptyChannelWithLength(t *testing.T) {
	// Test with an empty channel
	emptyCh := NewCharVarWithLength[int](5)
	emptyCh.StopWriting() // Close the empty channel
	_, emptyOK := emptyCh.ReadAllAvailableValues()

	// Check if reading from an empty channel returns false
	if !emptyOK {
		t.Error("Expected ReadAllAvailableValues to return true for an empty channel")
	}
}

func TestReadAllValues(t *testing.T) {
	// Create a new ChanVar instance
	ch := NewCharVar[int]()

	// Define values to be written to the channel
	values := []int{1, 2, 3}

	// Write values to the channel in a goroutine
	go func() {
		ch.WriteAllValue(values)
		ch.StopWriting()
	}()

	// Read all values from the channel
	readValues := ch.ReadAllValues()

	// Check if the number of read values matches the expected number
	if len(readValues) != len(values) {
		t.Errorf("Expected %d values, but got %d", len(values), len(readValues))
	}

	// Compare each read value with the expected value
	for i, v := range readValues {
		if v != values[i] {
			t.Errorf("Value mismatch at index %d, expected: %d, got: %d", i, values[i], v)
		}
	}
}

func TestReadAllValuesWithTimeout(t *testing.T) {
	// Create a new ChanVar instance
	ch := NewCharVar[int]()

	// Define values to be written to the channel
	values := []int{1, 2, 3}
	expectedValue := values[:2]

	// Write values to the channel in a goroutine
	go func() {
		for _, v := range values {
			ch.WriteValue(v)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Read all values from the channel with a timeout
	timeout := 200 * time.Millisecond
	readValues, ok := ch.ReadAllValuesWithTimeout(timeout)

	// Check if all values were read successfully
	if !ok {
		t.Error("Failed to read all values within the timeout")
	}

	// Check if the number of read values matches the expected number
	if len(readValues) != len(expectedValue) {
		t.Errorf("Expected %d values, but got %d", len(values), len(expectedValue))
	}

	// Compare each read value with the expected value
	for i, v := range expectedValue {
		if v != values[i] {
			t.Errorf("Value mismatch at index %d, expected: %d, got: %d", i, values[i], v)
		}
	}

	// Test timeout scenario
	emptyCh := NewCharVar[int]()
	timeout = 50 * time.Millisecond
	_, timeoutOK := emptyCh.ReadAllValuesWithTimeout(timeout)

	// Check if reading with timeout from an empty channel returns false
	if !timeoutOK {
		t.Error("Expected ReadAllValuesWithTimeout to return true due to timeout")
	}
}
