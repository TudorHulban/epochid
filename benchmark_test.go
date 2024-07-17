package epochid

import "testing"

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDRandom-16    	 2829472	       409.6 ns/op	      39 B/op	       2 allocs/op
func BenchmarkNewIDRandom(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDRandom()
	}
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDIncremental-16    	 6870781	       175.3 ns/op	      47 B/op	       3 allocs/op
func BenchmarkNewIDIncremental(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDIncremental10K()
	}
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDIncrementalWCorrection-16    	 6857497	       174.5 ns/op	      47 B/op	       3 allocs/op
func BenchmarkNewIDIncrementalWCorrection(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDIncremental10KWConCorrection()
	}
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDSimple-16    	 6545791	       179.5 ns/op	      47 B/op	       3 allocs/op
func BenchmarkNewIDSimple(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDIncrementalWConCorrection()
	}
}
