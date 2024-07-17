package epochid

import (
	"strconv"
	"time"
)

type EpochID int64

// NewEpochID provides correction in case container ID
// cannot be fetched. When found, container ID is 4 digits.
//
// Ignores parse error when converting to int64 for a faster result.
// Provides a sequence in the last 4 digits.
func NewEpochID() EpochID {
	containerID := getContainerID(4)

	now := strconv.FormatInt(time.Now().UnixNano(), 10)[:11]

	if len(containerID) == 0 {
		parsed, _ := strconv.ParseInt((now + getSequenceID()), 10, 64)

		return EpochID(parsed)
	}

	parsed, _ := strconv.ParseInt((now + containerID + getSequenceID()), 10, 64)

	return EpochID(parsed)
}

// Errors ignored for faster result.
func (e EpochID) GetUnixTimeSeconds() int64 {
	result, _ := strconv.ParseInt(
		strconv.FormatInt(int64(e), 10)[:11],
		10,
		64,
	)

	return result
}

func (e EpochID) String() string {
	return strconv.FormatInt(
		int64(e),
		10,
	)
}
