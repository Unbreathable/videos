package mutexbench_test

import (
	"mutexbench"
	"testing"
	"time"
)

func benchmarkSleeper(b *testing.B, s interface {
	Init()
	Run(time.Duration)
	Get() int64
}, d time.Duration) {
	s.Init()
	b.SetParallelism(WorkerCount)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.Run(d)
		}
	})
	b.ReportMetric(float64(s.Get()), "ops")
}

func BenchmarkSleeperMutex100ns(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, 100*time.Nanosecond)
}
func BenchmarkSleeperMutex250ns(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, 250*time.Nanosecond)
}
func BenchmarkSleeperMutex500ns(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, 500*time.Nanosecond)
}
func BenchmarkSleeprMutex500micros(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, 500*time.Microsecond)
}
func BenchmarkSleeperMutex1ms(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, time.Millisecond)
}
func BenchmarkSleeperMutex10ms(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, 10*time.Millisecond)
}
func BenchmarkSleeperMutex50ms(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperMutex{}, 50*time.Millisecond)
}

func BenchmarkSleeperWorker100ns(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, 100*time.Nanosecond)
}
func BenchmarkSleeperWorker250ns(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, 250*time.Nanosecond)
}
func BenchmarkSleeperWorker500ns(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, 500*time.Nanosecond)
}
func BenchmarkSleeperWorker500micros(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, 500*time.Microsecond)
}
func BenchmarkSleeperWorker1ms(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, time.Millisecond)
}
func BenchmarkSleeperWorker10ms(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, 10*time.Millisecond)
}
func BenchmarkSleeperWorker50ms(b *testing.B) {
	benchmarkSleeper(b, &mutexbench.SleeperWorker{}, 50*time.Millisecond)
}
