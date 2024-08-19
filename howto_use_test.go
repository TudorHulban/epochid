package epochid_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/TudorHulban/epochid"
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
