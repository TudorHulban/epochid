package epochid

import (
	"strconv"
)

type EpochID int64

func NewEpochID(item string) (EpochID, error) {
	var result EpochID

	if errValidation := validationString(
		&paramsValidationString{
			Caller: "NewEpochID",
			Name:   "item",
			Value:  item,
		},
		&result,
	); errValidation != nil {
		return 0,
			errValidation
	}

	return result,
		nil
}

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

// Formats as 1728-3090622-0000-3522.
func (e EpochID) UUIDFormat() string {
	asNumber := strconv.FormatInt(
		int64(e),
		10,
	)

	return asNumber[:4] + "-" + asNumber[4:11] + "-" + asNumber[11:15] + "-" + asNumber[15:]
}
