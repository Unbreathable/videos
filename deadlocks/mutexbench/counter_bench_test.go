package mutexbench_test

import (
	"mutexbench"
	"testing"
)

func BenchmarkCounterMutex(b *testing.B) {
	counter := &mutexbench.CounterMutex{}
	counter.Init()

	b.ResetTimer()
	for b.Loop() {
		counter.Increment()
	}
	b.ReportMetric(float64(counter.Get()), "ops")
}

func BenchmarkCounterMutexParallel(b *testing.B) {
	counter := &mutexbench.CounterMutex{}
	counter.Init()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
	b.ReportMetric(float64(counter.Get()), "ops")
}

func BenchmarkCounterAtomic(b *testing.B) {
	counter := &mutexbench.CounterAtomic{}
	counter.Init()

	b.ResetTimer()
	for b.Loop() {
		counter.Increment()
	}
	b.ReportMetric(float64(counter.Get()), "ops")
}

func BenchmarkCounterAtomicParallel(b *testing.B) {
	counter := &mutexbench.CounterAtomic{}
	counter.Init()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
	b.ReportMetric(float64(counter.Get()), "ops")
}

func BenchmarkCounterWorker(b *testing.B) {
	counter := &mutexbench.CounterWorker{}
	counter.Init()

	b.ResetTimer()
	for b.Loop() {
		counter.Increment()
	}
	b.ReportMetric(float64(counter.Get()), "ops")
}

func BenchmarkCounterWorkerParallel(b *testing.B) {
	counter := &mutexbench.CounterWorker{}
	counter.Init()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
	b.ReportMetric(float64(counter.Get()), "ops")
}
