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

func TestGetSequenceID(t *testing.T) {
	var res []string

	_sequenceID = 0

	n := 10
	for i := 0; i < n; i++ {
		res = append(res, getSequenceID())
	}

	want := []string{"0000", "0001", "0002", "0003", "0004", "0005", "0006", "0007", "0008", "0009"}

	require.Equal(t, want, res)
}

func TestSequenceReset(t *testing.T) {
	_sequenceID = 0

	for i := 0; i < 10000; i++ {
		getSequenceID()
	}

	require.Equal(t, "0000", getSequenceID())
}
