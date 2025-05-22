package epochid

import (
	"testing"
)

// ubuntu
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDSimple-16    	 6545791	       179.5 ns/op	      47 B/op	       3 allocs/op

// alma linux
// with all values precomputed
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDSimple-16    	10265635	       115.8 ns/op	      24 B/op	       1 allocs/op

// cpu: Apple M1
// BenchmarkNewIDSimple-8   	11679970	       104.1 ns/op	      24 B/op	       1 allocs/op
func BenchmarkNewIDSimple(b *testing.B) {
	generator := NewEpochGenerator()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		generator.GetValue()
	}
}
