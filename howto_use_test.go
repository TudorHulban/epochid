package epochid_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/TudorHulban/epochid"
)

func TestHowToUse(t *testing.T) {
	nowSeconds := time.Now().Unix()

	generator := epochid.NewEpochGenerator()

	id := generator.GetValue()

	fmt.Println("prefix:", nowSeconds)
	fmt.Println("ID (has prefix):", id)
}
