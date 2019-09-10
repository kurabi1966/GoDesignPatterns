package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 8
	expected := 13

	result := sum(a, b)

	if result != expected {
		t.Errorf("Our sum function does not work, %d+%d is not %d\n", a, b, result)
	}
}
