package epochid_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/TudorHulban/epochid"
)

func TestHowToUse(t *testing.T) {
	nowSeconds := time.Now().Unix()
	id := epochid.NewIDIncremental10KWConCorrection()

	fmt.Println("prefix:", nowSeconds)
	fmt.Println("ID (has prefix):", id)
}
