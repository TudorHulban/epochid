package epochid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	from1 := "a12b3"

	require.EqualValues(t,
		"190",
		hash(
			from1,
			3,
		),
		"1. happy path",
	)

	from2 := "abc"

	require.EqualValues(t,
		"298",
		hash(
			from2,
			3,
		),
		"2. no numbers",
	)

	from3 := ""

	require.EqualValues(t,
		"364",
		hash(
			from3,
			3,
		),
		"3. empty string passed",
	)
}

func BenchmarkPickNumbers(b *testing.B) {
	inputs := []struct {
		input   string
		howMany uint
	}{
		{"host123", 3},
		{"a1b2c3d4e5f6g7h8", 4},
		{"xyz", 2},
	}

	b.ResetTimer()

	for _, input := range inputs {
		b.Run(
			input.input,
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					hash(
						input.input,
						input.howMany,
					)
				}
			},
		)
	}
}
