package epochid_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/TudorHulban/epochid"
	"github.com/stretchr/testify/require"
)

func TestHowToUse(t *testing.T) {
	now := time.Now().UnixNano()

	generator := epochid.NewEpochGenerator()

	id := generator.GetValue()

	fmt.Println(
		"prefix:",
		strconv.FormatInt(now, 10)[:11],
	)
	fmt.Println("ID (has prefix):", id)
}

func TestNewEpochID(t *testing.T) {
	stringValue := "1737364079300002188"

	epochValue, errCr := epochid.NewEpochID(stringValue)
	require.NoError(t, errCr)
	require.Equal(t,
		strconv.Itoa(int(epochValue)),
		stringValue,
	)
}
