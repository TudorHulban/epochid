package epochid

import (
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStartValue(t *testing.T) {
	generator := NewEpochGenerator()

	epochID := generator.GetValue()

	idStr := strconv.FormatInt(int64(epochID), 10)

	sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

	require.Equal(t,
		"0000",
		sequencePart,
	)

	fmt.Println(
		t.Name(),
		epochID,
	)
}

func TestConcurrentGenerator(t *testing.T) {
	t.Log(t.Name())

	generator := NewEpochGenerator()

	size := 3

	var wg sync.WaitGroup
	var mu sync.Mutex
	var sequenceParts []string

	wg.Add(size)

	for ix := 0; ix < size; ix++ {
		go func() {
			epochID := generator.GetValue()

			idStr := strconv.FormatInt(int64(epochID), 10)

			sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

			mu.Lock()
			sequenceParts = append(sequenceParts, sequencePart)
			mu.Unlock()

			fmt.Println(
				t.Name(),
				sequencePart,
			)

			wg.Done()
		}()
	}

	wg.Wait()

	expected := []string{"0000", "0001", "0002"}

	require.ElementsMatch(t,
		expected,
		sequenceParts,
	)
}

func TestNotConcurrentGenerator(t *testing.T) {
	t.Log(t.Name())

	var sequenceParts []string

	generator := NewEpochGenerator()

	for range 5 {
		epochID := generator.GetValue()

		idStr := strconv.FormatInt(int64(epochID), 10)

		sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

		sequenceParts = append(sequenceParts, sequencePart)
	}

	expected := []string{"0000", "0001", "0002", "0003", "0004"}

	require.ElementsMatch(t,
		expected,
		sequenceParts,
	)
}

func TestSequenceLimitEpochGenerator(t *testing.T) {
	t.Log(t.Name())

	generator := NewEpochGenerator()

	// Consume a cycle of values.
	for i := 0; i < 9999; i++ {
		generator.GetValue()
	}

	fmt.Println(
		t.Name(),
		"10000th value",
		generator.GetValue(),
	)

	newCycleFirstEpochID := generator.GetValue()

	fmt.Println(
		t.Name(),
		"newCycleFirstEpochID:",
		newCycleFirstEpochID,
	)

	resetIDStr := strconv.FormatInt(int64(newCycleFirstEpochID), 10)

	resetSequencePart := resetIDStr[len(resetIDStr)-8 : len(resetIDStr)-4]
	if resetSequencePart != "0001" {
		t.Errorf(
			"1. Sequence Limit test failed: Expected sequence to reset to 0000, got %s, full ID: %s",
			resetSequencePart,
			resetIDStr,
		)
	}

	newCycleSecondEpochID := generator.GetValue()

	fmt.Println(
		t.Name(),
		"newCycleSecondEpochID:",
		newCycleSecondEpochID,
	)

	nextIDStr := strconv.FormatInt(int64(newCycleSecondEpochID), 10)

	nextSequencePart := nextIDStr[len(nextIDStr)-8 : len(nextIDStr)-4] // Corrected slicing
	if nextSequencePart != "0002" {
		t.Errorf(
			"2. Sequence Limit test failed: Expected sequence after reset to be 0001, got %s, full ID: %s",
			nextSequencePart,
			nextIDStr,
		)
	}
}

func TestEpochGeneratorOverall(t *testing.T) {
	t.Log(t.Name())

	generator := NewEpochGenerator()

	for i := 0; i < 5; i++ { // Test a few overall IDs
		epochID := generator.GetValue()

		idStr := strconv.FormatInt(int64(epochID), 10)
		if len(idStr) != 19 {
			t.Errorf(
				"Overall test failed: Generated EpochID is wrong: %s",
				idStr,
			)
		}

		_, errParse := strconv.ParseInt(idStr, 10, 64)
		if errParse != nil {
			t.Errorf(
				"Overall test failed: Generated EpochID is not a valid int64: %s, error: %v",
				idStr,
				errParse,
			)
		}

		sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

		_, errConversion := strconv.Atoi(sequencePart)
		if errConversion != nil {
			t.Errorf("Overall test failed: Sequence part of EpochID is not numeric: %s, error: %v",
				sequencePart,
				errConversion,
			)
		}
	}
}
