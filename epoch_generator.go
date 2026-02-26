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
	if gen.sequenceID.Load() >= _sequenceLimit-1 {
		gen.sequenceID.CompareAndSwap(_sequenceLimit-1, 0) // Ensures only one reset
	}

	return gen.precomputedIDs[gen.sequenceID.Add(1)-1]
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
