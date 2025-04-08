package epochid

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

func hash(text string, howMany uint) string {
	hasher := sha256.New()

	hasher.Write([]byte(text))
	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	if howMany > uint(len(hashString)) {
		howMany = uint(len(hashString))
	}

	numericHash, errParse := strconv.ParseUint(
		hashString[:howMany],
		16,
		64,
	)
	if errParse != nil {
		return ""
	}

	result := fmt.Sprintf("%d", numericHash)
	if len(result) > int(howMany) && howMany > 0 {
		return result[:howMany]
	}

	return result
}
