package ringbuffer

import "testing"

func BenchmarkRingBuffer(b *testing.B) {
	r := RingBuffer[int]{}

	var val int
	for i := 0; i < b.N; i++ {
		r.PushBack(val)
		r.PopFront()
		val++
	}
}
