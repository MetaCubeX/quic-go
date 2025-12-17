package ackhandler

// slicesAppendSeq appends the values from seq to the slice and
// returns the extended slice.
// If seq is empty, the result preserves the nilness of s.
func slicesAppendSeq[Slice ~[]E, E any](s Slice, seq func(yield func(E) bool)) Slice {
	seq(func(v E) bool {
		s = append(s, v)
		return true
	})
	return s
}

// slicesCollect collects values from seq into a new slice and returns it.
// If seq is empty, the result is nil.
func slicesCollect[E any](seq func(yield func(E) bool)) []E {
	return slicesAppendSeq([]E(nil), seq)
}
