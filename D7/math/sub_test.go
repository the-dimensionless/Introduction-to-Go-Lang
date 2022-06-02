package math

import (
	"testing"
)

func TestSub(t *testing.T) {

	got := Subtract(40, 10)
	want := 30

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
