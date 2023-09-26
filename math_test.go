package main

import "testing"

func TestSum(t *testing.T) {
	a := 10
	b := 5
	want := 15
	result := Sum(a, b)
	if result != want {
		t.Fatalf(`Sum(10, 5) = %d, want match for %d`, result, want)
	}
}

func TestSubtract(t *testing.T) {
	a := 101
	b := 100
	want := 1
	result := Subtract(a, b)

	if result != want {
		t.Fatalf(`Subtract(101, 100) = %d, want match for %d`, result, want)
	}
}

func TestMultiply(t *testing.T) {
	a := 10
	b := 12
	want := 120
	result := Multiply(a, b)

	if result != want {
		t.Fatalf(`Multiply(10, 12) = %d, want match for %d`, result, want)
	}
}

func TestDivide(t *testing.T) {
	a := 12
	b := 3
	want := 4
	result, err := Divide(a, b)

	if result != want || err != nil {
		t.Fatalf(`Divide(12, 3) = %d, %v, want match for %d`, result, err, want)
	}
}

func TestDivideWithZeroDivisor(t *testing.T) {
	a := 10
	b := 0
	want := 0
	result, err := Divide(a, b)

	if result != want || err == nil {
		t.Fatalf(`Divide(10, 0) = %d, %v, want match for %d, error`, result, err, want)
	}
}