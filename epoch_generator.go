package epochid

import (
	"strconv"
	"time"
)

type EpochGenerator struct {
	hostID string
}

// NewEpochGenerator provides correction in case container ID
// cannot be fetched. When found, container ID is 4 digits.
//
// Ignores parse error when converting to int64 for a faster result.
// Provides a sequence in the last 4 digits.
func NewEpochGenerator() *EpochGenerator {
	return &EpochGenerator{
		hostID: strconv.Itoa(int(getHostID(4))),
	}
}

func (gen *EpochGenerator) GetValue() EpochID {
	now := strconv.FormatInt(time.Now().UnixNano(), 10)[:11]

	if len(gen.hostID) == 0 {
		parsed, _ := strconv.ParseInt((now + getSequenceID()), 10, 64)

		return EpochID(parsed)
	}

	parsed, _ := strconv.ParseInt((now + gen.hostID + getSequenceID()), 10, 64)

	return EpochID(parsed)
}
