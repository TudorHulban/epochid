package epochid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	slice := []EpochID{
		1729954304000012180,
		1729954304000012181,
		1729954304000012182,
		1729954304000012185,
		1729954304000010001,
	}

	lookingFor := 1729954304000010001
	notLookingFor := 1729954304000010007

	require.True(t,
		Epochs(slice).Contains(
			lookingFor,
		),
	)

	require.False(t,
		Epochs(slice).Contains(
			notLookingFor,
		),
	)
}
