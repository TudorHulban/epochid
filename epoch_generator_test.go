package epochid

import (
	"strconv"
	"sync"
	"testing"
)

func TestConcurrentGenerator(t *testing.T) {
	generator := NewEpochGenerator()

	size := 2

	var wg sync.WaitGroup

	wg.Add(size)

	for ix := 0; ix < size; ix++ {
		go func() {
			generator.GetValue()

			wg.Done()
		}()
	}

	wg.Wait()
}

func TestPrecomputedEpochGenerator(t *testing.T) {
	generator := NewEpochGenerator()

	for i := 1; i <= _precomputedSize; i++ {
		epochID := generator.GetValue()

		idStr := strconv.FormatInt(int64(epochID), 10)

		sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

		expectedSequence := generator.precomputedIDs[i-1]
		if sequencePart != expectedSequence {
			t.Errorf(
				"Precomputed test failed at index %d: expected sequence %s, got %s, full ID %s",
				i-1,
				expectedSequence,
				sequencePart,
				idStr,
			)
		}
	}
}

func TestSequenceLimitEpochGenerator(t *testing.T) {
	generator := NewEpochGenerator()

	// Consume almost all 9999 sequence IDs
	for i := 0; i <= 9999; i++ {
		generator.GetValue()
	}

	next1EpochID := generator.GetValue()

	resetIDStr := strconv.FormatInt(int64(next1EpochID), 10)

	resetSequencePart := resetIDStr[len(resetIDStr)-8 : len(resetIDStr)-4]
	if resetSequencePart != "0000" {
		t.Errorf(
			"Sequence Limit test failed: Expected sequence to reset to 0000, got %s, full ID: %s",
			resetSequencePart,
			resetIDStr,
		)
	}

	next2EpochID := generator.GetValue()

	nextIDStr := strconv.FormatInt(int64(next2EpochID), 10)

	nextSequencePart := nextIDStr[len(nextIDStr)-8 : len(nextIDStr)-4] // Corrected slicing
	if nextSequencePart != "0001" {
		t.Errorf(
			"Sequence Limit test failed: Expected sequence after reset to be 0001, got %s, full ID: %s",
			nextSequencePart,
			nextIDStr,
		)
	}
}

func TestEpochGeneratorOverall(t *testing.T) {
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
