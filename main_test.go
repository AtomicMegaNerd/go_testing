package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got '%v', Wanted: '%v'", got, expected)
	}
}

// Tests must begin with the Test* prefix
func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 5
	if got != expected {
		t.Errorf("Did not get expected result. Got '%v', Wanted: '%v'", got, expected)
	}
}
