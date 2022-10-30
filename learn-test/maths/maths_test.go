package maths

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(5, 5)
	if sum != 10 {
		t.Errorf("Sum is incorrect! got: %d, want: %d\n", sum, 10)
	}
}

func TestSubtract(t *testing.T) {
	difference := Subtract(10, 5)
	if difference != 5 {
		t.Errorf("Difference is incorrect! got: %d, want: %d\n", difference, 5)
	}
}

func TestMultiply(t *testing.T) {
	product := Multiply(5, 5)
	if product != 25 {
		t.Errorf("Produt is incorrect! got: %d, want: %d\n", product, 25)
	}
}

func TestDivide(t *testing.T) {
	quotient := Divide(5, 5)
	if quotient != 1 {
		t.Errorf("Sum is incorrect! got: %d, want: %d\n", quotient, 1)
	}
}
