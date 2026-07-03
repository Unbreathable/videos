package mutexbench_test

import (
	"fmt"
	"mutexbench"
	"testing"
)

func benchmarkHasher(b *testing.B, s interface {
	Init()
	Run()
	Get() int64
}) {
	s.Init()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Run()
		}
	})
	b.ReportMetric(float64(s.Get()), "ops")
}

func BenchmarkHashingMutex(b *testing.B) {
	for n := 1; n <= 10; n++ {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			benchmarkHasher(b, &mutexbench.HasherMutex{N: n})
		})
	}
}

func BenchmarkHashingWorker(b *testing.B) {
	for n := 1; n <= 10; n++ {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			benchmarkHasher(b, &mutexbench.HasherWorker{N: n})
		})
	}
}
