//go:build !go1.25

package synctest

import (
	"testing"
)

func Test(t *testing.T, f func(t *testing.T)) {
	t.Skip("don't support synctest")
}

func Wait() {}
