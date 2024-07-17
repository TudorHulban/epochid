package epochid

import (
	"strconv"
	"time"
)

type EpochID int64

// NewIDIncrementalWConCorrection provides correction in case container ID
// cannot be fetched. When found container ID is 4 digits.
//
// Ignores parse error when converting to uint for a faster result.
// Provides a sequence in the last 4 digits.
func NewIDIncrementalWConCorrection() EpochID {
	containerID := getContainerID(4)

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	if len(containerID) == 0 {
		parsed, _ := strconv.ParseInt((now[:11] + getSequenceID()), 10, 64)

		return EpochID(parsed)
	}

	parsed, _ := strconv.ParseInt((now[:11] + containerID + getSequenceID()), 10, 64)

	return EpochID(parsed)
}

// Errors ignored for faster result.
func (e EpochID) GetUnixTimeSeconds() int64 {
	timestamp := strconv.FormatInt(int64(e), 10)

	result, _ := strconv.ParseInt(timestamp[:11], 10, 64)

	return result
}
