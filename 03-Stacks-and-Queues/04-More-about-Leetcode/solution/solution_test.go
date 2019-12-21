package solution

import "testing"

func TestIsValid(t *testing.T) {
	assertValidString := func(t *testing.T, str string, want bool) {
		t.Helper()

		got := IsValid(str)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	assertValidString(t, "()[]{}", true)
	assertValidString(t, "([)]", false)
}
