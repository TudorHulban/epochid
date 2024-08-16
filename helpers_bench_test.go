package epochid

import "testing"

func BenchmarkRandInt(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randInt()
	}
}
