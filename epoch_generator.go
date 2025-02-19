package epochid

import (
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const _precomputedSize = 256

var _sequenceID atomic.Int64

type EpochGenerator struct {
	hostID         string
	precomputedIDs [_precomputedSize]string
}

// NewEpochGenerator provides correction in case container ID
// cannot be fetched. When found, container ID is 4 digits.
//
// Ignores parse error when converting to int64 for a faster result.
// Provides a sequence in the last 4 digits.
func NewEpochGenerator() *EpochGenerator {
	result := EpochGenerator{
		hostID: strconv.Itoa(int(getHostID(4))),
	}

	result.initializePrecomputedIDs()

	return &result
}

func (gen *EpochGenerator) initializePrecomputedIDs() {
	for i := 0; i < _precomputedSize; i++ {
		gen.precomputedIDs[i] = strings.Repeat(
			"0",
			4-len(strconv.FormatInt(int64(i), 10))) + strconv.FormatInt(int64(i), 10)
	}
}

func (gen *EpochGenerator) GetValue() EpochID {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)[:11]

	if len(gen.hostID) == 0 {
		parsed, _ := strconv.ParseInt((now + gen.getSequenceID()), 10, 64)

		return EpochID(parsed)
	}

	parsed, _ := strconv.ParseInt((now + gen.getSequenceID() + gen.hostID), 10, 64)

	return EpochID(parsed)
}

func (gen *EpochGenerator) getSequenceID() string {
	current := _sequenceID.Add(1)

	if current < _precomputedSize {
		return gen.precomputedIDs[current-1]
	}

	if current == 10000 {
		_sequenceID.Store(0)

		return gen.precomputedIDs[0]
	}

	s := strconv.FormatInt(current-1, 10)

	return strings.Repeat("0", 4-len(s)) + s
}
