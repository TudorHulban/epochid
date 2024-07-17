package epochid

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleNewID(t *testing.T) {
	id := NewEpochID()

	fmt.Printf(
		"%s\n%s\n%d\nunix seconds:%d\n",
		t.Name(),
		id,
		time.Now().UnixNano(),
		id.GetUnixTimeSeconds(),
	)
}

func TestStringerSimpleNewID(t *testing.T) {
	p := func(value fmt.Stringer) {
		fmt.Printf(
			"%s\n",
			value.String(),
		)
	}

	id := NewEpochID()

	p(id)
}
