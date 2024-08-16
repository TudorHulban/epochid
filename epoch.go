package epochid

import (
	"strconv"
)

type EpochID int64

// Errors ignored for faster result.
func (e EpochID) GetUnixTimeSeconds() int64 {
	result, _ := strconv.ParseInt(
		strconv.FormatInt(int64(e), 10)[:10],
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
