package calculator

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(2, 1) != 1 {
		t.Error("2 - 1 did not equal 1")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 3) != 6 {
		t.Error("2 * 3 did not equal 6")
	}
}

func TestDivide(t *testing.T) {
	if Divide(6, 2) != 3 {
		t.Error("6 / 2 did not equal 3")
	}
}

func TestRemainder(t *testing.T) {

	if Remainder(5, 2) != 1 {
		t.Error("5 % 2 did not equal 1")
	}
}

func TestSquare(t *testing.T) {
	if Square(2) != 4 {
		t.Error("2 * 2 did not equal 4")
	}
}

func TestCube(t *testing.T) {
	if Cube(2) != 8 {
		t.Error("2 * 2 * 2 did not equal 8")
	}
}
