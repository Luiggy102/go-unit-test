package sometests

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	result := Add(20, 2)
	expect := 22

	if result != expect {
		t.Errorf("got %d, exppected %d", result, expect)
	}
}
