package logk

import "testing"

func BenchmarkDebug(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Debug("benchmark testing...")
	}
}
