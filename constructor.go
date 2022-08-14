package epochid

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var _containerID string
var _sequenceID int64

func getContainerID() string {
	if len(_containerID) == 0 {
		id, err := readContainerID()
		if err != nil {
			return ""
		}

		_containerID = pickNumbersFrom(id, 5)
	}

	return _containerID
}

func getSequenceID() string {
	var mu sync.Mutex
	var id int64

	mu.Lock()
	id = _sequenceID

	_sequenceID++

	if _sequenceID == 10000 {
		_sequenceID = 0
	}

	mu.Unlock()

	res := "000" + strconv.FormatInt(id, 10)

	return res[len(res)-4:]
}

func NewIDRandom() (uint64, error) {
	containerID := getContainerID()
	if len(containerID) == 0 {
		return 0, errors.New("there was an error")
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	random, errRa := randInt()
	if errRa != nil {
		return 0, fmt.Errorf("NewID strconv.ParseUint: %w", errRa)
	}

	parsed, errPa := strconv.ParseUint((now[:11] + containerID + random + now[16:19])[:20], 10, 64)
	if errPa != nil {
		return 0, fmt.Errorf("NewID strconv.ParseUint: %w", errPa)
	}

	return parsed, nil
}

// NewIDIncremental10K provides a sequence in the last 4 digits.
func NewIDIncremental10K() (uint64, error) {
	containerID := getContainerID()
	if len(containerID) == 0 {
		return 0, errors.New("there was an error")
	}

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	parsed, errPa := strconv.ParseUint((now[:11] + containerID + getSequenceID()), 10, 64)
	if errPa != nil {
		return 0, fmt.Errorf("NewID strconv.ParseUint: %w", errPa)
	}

	return parsed, nil
}

// NewIDIncremental10KWCoCorrection provides correction in case container ID cannot be fetched.
// Ignores parse error when converting to uint for a faster result.
// Provides a sequence in the last 4 digits.
func NewIDIncremental10KWCoCorrection() uint64 {
	containerID := getContainerID()

	now := strconv.FormatInt(time.Now().UnixNano(), 10)

	if len(containerID) == 0 {
		parsed, _ := strconv.ParseUint((now[:16] + getSequenceID()), 10, 64)

		return parsed
	}

	parsed, _ := strconv.ParseUint((now[:11] + containerID + getSequenceID()), 10, 64)

	return parsed
}
