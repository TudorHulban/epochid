package epochid_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/TudorHulban/epochid"
)

func TestHowToUse(t *testing.T) {
	now := time.Now().Unix()
	id := epochid.NewIDIncremental10KWCoCorrection()

	fmt.Println("prefix:", now)
	fmt.Println("ID (has prefix):", id)
}
