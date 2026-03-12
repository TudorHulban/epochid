package epochid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandIntID(t *testing.T) {
	value1, errID := randInt()
	require.NoError(t, errID)
	require.GreaterOrEqual(t, len(value1), 1)

	value2, _ := randInt()
	require.GreaterOrEqual(t, len(value2), 1)
}
