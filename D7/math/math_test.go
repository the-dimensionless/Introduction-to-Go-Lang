package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSomething(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello2"

	assert.Equal(t, a, b, "The two words should be the same.")

}
