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
		"0001",
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

	var (
		wg            sync.WaitGroup
		mu            sync.Mutex
		sequenceParts []string
	)

	wg.Add(size)

	for range size {
		go func() {
			epochID := generator.GetValue()

			idStr := strconv.FormatInt(int64(epochID), 10)

			sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

			mu.Lock()
			sequenceParts = append(sequenceParts, sequencePart) //nolint:wsl_v5
			mu.Unlock()

			fmt.Println(
				t.Name(),
				sequencePart,
			)

			wg.Done()
		}()
	}

	wg.Wait()

	expected := []string{"0001", "0002", "0003"}

	require.ElementsMatch(t,
		expected,
		sequenceParts,
	)
}

func TestNotConcurrentGenerator(t *testing.T) {
	t.Log(t.Name())

	upTo := 5

	sequenceParts := make([]string, upTo)

	generator := NewEpochGenerator()

	for ix := range upTo {
		epochID := generator.GetValue()

		idStr := strconv.FormatInt(int64(epochID), 10)

		sequencePart := idStr[len(idStr)-8 : len(idStr)-4]

		sequenceParts[ix] = sequencePart
	}

	expected := []string{"0001", "0002", "0003", "0004", "0005"}

	require.ElementsMatch(t,
		expected,
		sequenceParts,
	)
}

func TestSequenceLimitEpochGenerator(t *testing.T) {
	t.Log(t.Name())

	generator := NewEpochGenerator()

	// Consume a cycle of values.
	for range 9999 {
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

	for range 5 { // Test a few overall IDs
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

func TestEpochGenerator_RolloverRace(t *testing.T) {
	t.Log(t.Name())

	gen := NewEpochGenerator()
	gen.sequenceID.Store(_sequenceLimit - 1)

	const (
		workers    = 64
		iterations = 128
	)

	var wg sync.WaitGroup
	wg.Add(workers)

	var (
		mu       sync.Mutex
		seenAt   = make(map[string]int) // seq -> first goroutine+call that saw it
		failures []string
	)

	for w := range workers {
		go func(worker int) {
			defer wg.Done()

			for i := range iterations {
				seq := gen.getSequenceID()

				mu.Lock()
				if prev, exists := seenAt[seq]; exists {
					failures = append(failures,
						fmt.Sprintf(
							"worker %d iter %d: duplicate %s (first seen at call %d)",
							w,
							i,
							seq,
							prev,
						),
					)
				} else {
					seenAt[seq] = w*iterations + i
				}
				mu.Unlock()
			}
		}(w)
	}

	wg.Wait()

	for _, f := range failures {
		t.Error(f)
	}
}
