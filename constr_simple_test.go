package epochid

import (
	"fmt"
	"testing"
	"time"
)

func TestSimpleNewID(t *testing.T) {
	id := NewEpochID()

	// id := EpochID(1721634797482180004)

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
