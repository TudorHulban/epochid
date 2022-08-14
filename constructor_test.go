package epochid

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIDRandom(t *testing.T) {
	id, errNew := NewIDRandom()
	require.NoError(t, errNew)

	fmt.Println("TestNewIDRandom example ID:", id)
}

func TestNewIDIncremental10K(t *testing.T) {
	id1, errNew := NewIDIncremental10K()
	require.NoError(t, errNew)

	id2, _ := NewIDIncremental10K()
	require.Equal(t, uint64(1), id2-id1, fmt.Sprintf("id1: %d, id2: %d", id1, id2))
}

func TestNewIDIncremental10KWCoCorrection(t *testing.T) {
	id1 := NewIDIncremental10KWCoCorrection()

	var wg sync.WaitGroup
	wg.Add(1)

	var id2 uint64

	go func() {
		id2 = NewIDIncremental10KWCoCorrection()

		wg.Done()
	}()

	wg.Wait()

	require.Equal(t, uint64(1), id2-id1, fmt.Sprintf("id1: %d, id2: %d", id1, id2))
}
