package main

import "testing"

func TestCalc(t *testing.T) {
	got := Calc(10, 10)
	if got != 100 {
		t.Errorf("got: %q, expected: 100", got)
	}
}
