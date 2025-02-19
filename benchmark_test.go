package epochid

import (
	"testing"
)

// ubuntu - performance
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDSimple-16    	 6545791	       179.5 ns/op	      47 B/op	       3 allocs/op

// rocky - balanced
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDSimple-16    	 4493860	       725.5 ns/op	      39 B/op	       2 allocs/op

// alma linux - balanced
// with all values precomputed
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDSimple-16    	 9781618	       118.4 ns/op	      24 B/op	       1 allocs/op
func BenchmarkNewIDSimple(b *testing.B) {
	generator := NewEpochGenerator()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generator.GetValue()
	}
}
