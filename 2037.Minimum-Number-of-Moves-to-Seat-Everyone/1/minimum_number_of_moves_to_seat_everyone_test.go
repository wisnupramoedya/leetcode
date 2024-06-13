package minimum_number_of_moves_to_seat_everyone

import "testing"

func Test2037_1_1(t *testing.T) {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}
	output := 4

	result := minMovesToSeat(seats, students)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test2037_1_2(t *testing.T) {
	seats := []int{4, 1, 5, 9}
	students := []int{1, 3, 2, 6}
	output := 7

	result := minMovesToSeat(seats, students)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test2037_1_3(t *testing.T) {
	seats := []int{2, 2, 6, 6}
	students := []int{1, 3, 2, 6}
	output := 4

	result := minMovesToSeat(seats, students)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}

func Test2037_1_4(t *testing.T) {
	seats := []int{12, 14, 19, 19, 12}
	students := []int{19, 2, 17, 20, 7}
	// 0, 10, 3, 1, 5
	output := 19

	result := minMovesToSeat(seats, students)

	if output != result {
		t.Errorf("Expected %d, but got %d", output, result)
	}
}
