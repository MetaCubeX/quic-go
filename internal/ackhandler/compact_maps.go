package ackhandler

// mapsInsert adds the key-value pairs from seq to m.
// If a key in seq already exists in m, its value will be overwritten.
func mapsInsert[Map ~map[K]V, K comparable, V any](m Map, seq func(yield func(K, V) bool)) {
	seq(func(k K, v V) bool {
		m[k] = v
		return true
	})
}

// Collect collects key-value pairs from seq into a new map
// and returns it.
func mapsCollect[K comparable, V any](seq func(yield func(K, V) bool)) map[K]V {
	m := make(map[K]V)
	mapsInsert(m, seq)
	return m
}
