package epochid

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func NewIDRandom() (uint64, error) {
	containerID := getContainerID(5)
	if len(containerID) == 0 {
		return 0,
			errors.New("could not get container ID")
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	random, errRa := randInt()
	if errRa != nil {
		return 0,
			fmt.Errorf("NewIDRandom strconv.ParseUint: %w", errRa)
	}

	parsed, errParse := strconv.ParseUint((now[:11] + containerID + random + now[16:19])[:20], 10, 64)
	if errParse != nil {
		return 0,
			fmt.Errorf("NewIDRandom strconv.ParseUint: %w", errParse)
	}

	return parsed, nil
}

// NewIDIncremental10K provides a sequence in the last 4 digits.
func NewIDIncremental10K() (uint64, error) {
	containerID := getContainerID(5)
	if len(containerID) == 0 {
		return 0,
			errors.New("could not get container ID")
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	parsed, errPa := strconv.ParseUint((now[:11] + containerID + getSequenceID()), 10, 64)
	if errPa != nil {
		return 0,
			fmt.Errorf("NewID strconv.ParseUint: %w", errPa)
	}

	return parsed, nil
}

// NewIDIncremental10KWConCorrection provides correction in case container ID cannot be fetched.
// Ignores parse error when converting to uint for a faster result.
// Provides a sequence in the last 4 digits.
func NewIDIncremental10KWConCorrection() uint64 {
	containerID := getContainerID(5)

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	if len(containerID) == 0 {
		parsed, _ := strconv.ParseUint((now[:16] + getSequenceID()), 10, 64)

		return parsed
	}

	parsed, _ := strconv.ParseUint((now[:11] + containerID + getSequenceID()), 10, 64)

	return parsed
}
