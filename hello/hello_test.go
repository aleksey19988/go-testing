package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("")
	want := "Hello, World!"

	if got != want {
		t.Errorf("Hello(\"\") = %q, want %q", got, want)
	}

	got = Hello("Denny")
	want = "Hello, Denny!"

	if got != want {
		t.Errorf("Hello(\"Denny\") = %q, want %q", got, want)
	}
}
