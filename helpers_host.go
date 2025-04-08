package epochid

import (
	"fmt"
	"os"
)

func readContainerID() (string, error) {
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

func getHostID(length uint) string {
	id, errHostID := readContainerID()
	if errHostID != nil {
		return ""
	}

	return hash(id, length)
}
