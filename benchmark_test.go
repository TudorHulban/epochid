package linuxid

import "testing"

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkConstructor-12    	 1896780	       794.5 ns/op	      71 B/op	       3 allocs/op
func BenchmarkNewIDRandom(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDRandom()
	}
}

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkConstructor-12    	 4363184	       349.8 ns/op	      71 B/op	       4 allocs/op
func BenchmarkNewIDIncremental(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDIncremental10K()
	}
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkNewIDIncrementalWCorrection-16    	 3401972	       386.4 ns/op	      71 B/op	       4 allocs/op
func BenchmarkNewIDIncrementalWCorrection(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NewIDIncremental10KWCoCorrection()
	}
}
