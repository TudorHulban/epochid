package epochid

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

// randInt generates a random uint32 as 4 digit string
// time cost is around 500 ns.
func randInt() (string, error) {
	buf := make([]byte, 3)

	if _, err := rand.Reader.Read(buf); err != nil {
		return "",
			fmt.Errorf("generate random number: %w;", err)
	}

	res := uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])

	return strconv.Itoa(int(res)),
		nil
}
