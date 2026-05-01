package even

import "testing"

func TestIsEven(t *testing.T) {
	got := IsEven(1)
	want := false

	if got != want {
		t.Errorf("IsEven() = %v; want %v", got, want)
	}

	got = IsEven(2)
	want = true

	if got != want {
		t.Errorf("IsEven() = %v; want %v", got, want)
	}
}
