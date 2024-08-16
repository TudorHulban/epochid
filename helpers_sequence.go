package epochid

import (
	"strconv"
	"sync/atomic"
)

var _sequenceID int64

func getSequenceID() string {
	result := "000" + strconv.FormatInt(
		atomic.LoadInt64(&_sequenceID),
		10,
	)

	if atomic.LoadInt64(&_sequenceID) == 10000 {
		atomic.StoreInt64(&_sequenceID, 0)
	}

	atomic.AddInt64(
		&_sequenceID,
		1,
	)

	return result[len(result)-4:]
}
