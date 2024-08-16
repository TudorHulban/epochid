package epochid

import (
	"fmt"
	"os"
)

// readHostID returns machine ID
func readHostID() (string, error) {
	idHardware, errID := os.ReadFile("/etc/machine-id")
	if errID != nil || len(idHardware) == 0 {
		hostname, errHo := os.Hostname()
		if errHo != nil {
			return "",
				fmt.Errorf(
					"readContainerID os.Hostname: %w",
					errHo,
				)
		}

		return hostname, nil
	}

	return string(idHardware), nil
}

func getHostID(length uint) int64 {
	id, errHostID := readHostID()
	if errHostID != nil {
		return 0
	}

	return pickNumbersFrom(id, length)
}