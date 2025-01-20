package epochid

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSimpleNewID(t *testing.T) {
	id := NewEpochGenerator().GetValue()

	fmt.Printf(
		"%s\n%s\n%d\nunix seconds:%d\n",
		t.Name(),
		id,
		time.Now().UnixNano(),
		id.GetUnixTimeSeconds(),
	)
	fmt.Println(
		time.Unix(id.GetUnixTimeSeconds(), 0),
	)

	fmt.Println(id)
	fmt.Println(id.UUIDFormat())
}

func TestStringerSimpleNewID(t *testing.T) {
	p := func(value fmt.Stringer) {
		fmt.Printf(
			"%s\n",
			value.String(),
		)
	}

	id := NewEpochGenerator().GetValue()

	p(id)
}

func TestErrorsNewEpochID(t *testing.T) {
	stringValue := "173736407930000218"

	epochValue, errCr := NewEpochID(stringValue)
	require.Error(t, errCr)
	require.Zero(t, epochValue)
}
