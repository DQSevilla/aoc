package main

import "testing"

func testDepths() []int {
	return []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
}

func TestPartOne(t *testing.T) {
	want, got := 7, partOne(testDepths())

	if want != got {
		t.Fatalf("Want %d Got %d", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	want, got := 5, partTwo(testDepths())

	if want != got {
		t.Fatalf("Want %d Got %d", want, got)
	}
}
