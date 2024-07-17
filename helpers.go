package epochid

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func pickNumbersFrom(s string, howMany uint) string {
	isNumber := func(letter string) bool {
		numbers := []string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
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
			return strings.Join(res, "")
		}
	}

	return ""
}

// randInt generates a random uint32 as 4 digit string
// time cost is around 500 ns.
func randInt() (string, error) {
	buf := make([]byte, 3)

	if _, err := rand.Reader.Read(buf); err != nil {
		return "",
			fmt.Errorf("generate random number: %W;", err)
	}

	res := uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])

	return strconv.Itoa(int(res)), nil
}

// readMachineID returns machine ID
func readContainerID() (string, error) {
	idHardware, errID := os.ReadFile("/etc/machine-id")
	if errID != nil || len(idHardware) == 0 {
		hostname, errHo := os.Hostname()
		if errHo != nil {
			return "",
				fmt.Errorf("readContainerID os.Hostname: %w", errHo)
		}

		return hostname, nil
	}

	return string(idHardware), nil
}

func getContainerID(length uint) string {
	if len(_containerID) == 0 {
		id, err := readContainerID()
		if err != nil {
			return ""
		}

		_containerID = pickNumbersFrom(id, length)
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
