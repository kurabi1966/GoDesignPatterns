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

func TestMultiply(t *testing.T) {
	a := 5
	b := 8
	expected := 40

	result := multiply(a, b)
	if result != expected {
		t.Errorf("Our multiply function does not work correctly, %d x %d is not %d", a, b, result)
	}
}
