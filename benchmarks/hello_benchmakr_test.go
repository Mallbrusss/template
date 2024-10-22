package benchmarks

import (
	"testing"
	"internal"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		internal.HelloWorld()
	}
}
