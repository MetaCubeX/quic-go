//go:build go1.25

package synctest

import (
	"testing"
	"testing/synctest"
)

func Test(t *testing.T, f func(t *testing.T)) {
	t.Setenv("GODEBUG", "asynctimerchan=0")
	synctest.Test(t, f)
}

func Wait() {
	synctest.Wait()
}
