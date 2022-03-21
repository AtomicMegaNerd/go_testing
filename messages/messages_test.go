package messages

// package messages_test (use this for black box testing)

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Failed. Expected %q, Got %q", got, expect)
	}
}

func TestGreetTableDrive(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Called Greet() with Input: %v, Expected: %q, Got: %q", s.input, s.expect, got)
		}
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"

	if got != expect {
		t.Errorf("Failed. Expected: %q Got: %q", got, expect)
	}
}
