package epochid

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZeroHostIDNewEpochGenerator(t *testing.T) {
	generator := NewEpochGenerator()

	value := generator.GetValue()

	require.NotZero(t, value)
	require.Len(t,
		strconv.Itoa(int(value)),
		19,
	)
}
