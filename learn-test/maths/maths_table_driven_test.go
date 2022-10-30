package maths

import "testing"

func TestAddTD(t *testing.T) {
	tbl := []struct {
		num1 int
		num2 int
		want int
	}{
		{5, 5, 10},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -1, -2},
	}

	for _, testCase := range tbl {
		output := Add(testCase.num1, testCase.num2)
		if output != testCase.want {
			t.Errorf("Sum is incorrect! got: %d, want: %d\n", output, testCase.want)
		}
	}
}

func TestSubtractTD(t *testing.T) {

	tbl := []struct {
		num1 int
		num2 int
		want int
	}{
		{5, 5, 0},
		{0, 0, 0},
		{-1, 1, -2},
		{-1, -1, 0},
	}

	for _, testCase := range tbl {
		output := Subtract(testCase.num1, testCase.num2)
		if output != testCase.want {
			t.Errorf("Difference is incorrect! got: %d, want: %d\n", output, testCase.want)
		}
	}

}

func TestMultiplyTD(t *testing.T) {
	tbl := []struct {
		num1 int
		num2 int
		want int
	}{
		{5, 5, 25},
		{0, 0, 0},
		{-1, 1, -1},
		{-1, -1, 1},
	}

	for _, testCase := range tbl {
		output := Multiply(testCase.num1, testCase.num2)
		if output != testCase.want {
			t.Errorf("Product is incorrect! got: %d, want: %d\n", output, testCase.want)
		}
	}
}

func TestDivideTD(t *testing.T) {

	tbl := []struct {
		num1 int
		num2 int
		want int
	}{
		{5, 5, 1},
		{15, 3, 5},
		{-1, 1, -1},
		{-1, -1, 1},
	}

	for _, testCase := range tbl {
		output := Divide(testCase.num1, testCase.num2)
		if output != testCase.want {
			t.Errorf("Quotient is incorrect! got: %d, want: %d\n", output, testCase.want)
		}
	}

}
