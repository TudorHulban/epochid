package epochid

import (
	"sync"
	"testing"
)

func TestGeneratorLoad(t *testing.T) {
	generator := NewEpochGenerator()

	size := 2

	var wg sync.WaitGroup

	wg.Add(size)

	for ix := 0; ix < size; ix++ {
		go func() {
			generator.GetValue()

			wg.Done()
		}()
	}

	wg.Wait()
}
