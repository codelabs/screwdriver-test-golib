package hello

import (
	"strings"
	"testing"
)

func TestGreetings(t *testing.T) {
	var expected = "Hello Go !! from screwdriver"
	var received = greetings()

	if !strings.EqualFold(expected, received) {
		t.Errorf("result == %q, want %q", received, expected)
	}
}
