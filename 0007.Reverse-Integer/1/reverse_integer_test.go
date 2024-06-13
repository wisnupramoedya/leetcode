package reverse_integer

import "testing"

func Test7_1_1(t *testing.T) {
	x := 123
	output := 321

	result := reverse(x)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test7_1_2(t *testing.T) {
	x := -123
	output := -321

	result := reverse(x)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test7_1_3(t *testing.T) {
	x := 120
	output := 21

	result := reverse(x)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test7_1_4(t *testing.T) {
	x := 1563847412
	output := 0

	result := reverse(x)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test7_1_5(t *testing.T) {
	x := -1563847412
	output := 0

	result := reverse(x)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}
