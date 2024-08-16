package epochid

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

func pickNumbersFrom(s string, howMany uint) int64 {
	isNumber := func(letter string) bool {
		numbers := []string{
			"1", "2", "3", "4", "5", "6", "7", "8", "9",
		}

		for _, number := range numbers {
			if letter == number {
				return true
			}
		}

		return false
	}

	var j int
	res := make([]string, 0)

	for i := len(s); i >= 0; i-- {
		if isNumber(s[i-1 : i]) {
			res = append(res, s[i-1:i])

			j++
		}

		if j == int(howMany) {
			result, _ := strconv.Atoi(
				strings.Join(res, ""),
			)

			return int64(result)
		}
	}

	return 0
}

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
