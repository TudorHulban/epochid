package epochid

import (
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const _sequenceLimit = 10000

type EpochGenerator struct {
	hostID         string
	precomputedIDs [_sequenceLimit]string // from 0000 to 9999
	sequenceID     atomic.Int64
}

// NewEpochGenerator provides correction in case container ID
// cannot be fetched. When found, container ID is 4 digits.
//
// Ignores parse error when converting to int64 for a faster result.
// Provides a sequence in the last 4 digits.
func NewEpochGenerator() *EpochGenerator {
	result := EpochGenerator{
		hostID: getHostID(4),
	}

	result.initializePrecomputedIDs()

	return &result
}

func (gen *EpochGenerator) initializePrecomputedIDs() {
	for i := range _sequenceLimit {
		gen.precomputedIDs[i] = strings.Repeat(
			"0",
			4-len(strconv.FormatInt(int64(i), 10))) +
			strconv.FormatInt(int64(i), 10)
	}
}

func (gen *EpochGenerator) getSequenceID() string {
	// standard lock-free CAS loop pattern that
	// eliminates the time-of-check to time-of-use (TOCTOU) gap entirely.

	// Only one goroutine can successfully transition from current to next.
	// All other goroutines observing the same current will fail CAS and retry.
	// Rollover (next = 0) is validated inside the CAS, not after a separate increment.
	// No goroutine can ever increment past _sequenceLimit-1 because the CAS enforces the transition.
	for {
		current := gen.sequenceID.Load()

		next := current + 1
		if next >= _sequenceLimit {
			next = 0
		}

		if gen.sequenceID.CompareAndSwap(current, next) {
			return gen.precomputedIDs[next]
		}
	}
}

func (gen *EpochGenerator) GetValue() EpochID {
	parsed, _ := strconv.
		ParseInt(
			(strconv.FormatInt(time.Now().UnixNano(), 10)[:11] +
				gen.getSequenceID() + gen.hostID),
			10,
			64,
		)

	return EpochID(parsed)
}
