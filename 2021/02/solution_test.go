package main

import "testing"

func testCommands() []Command {
	return []Command{
		Command{Dir: Forward, Delta: 5},
		Command{Dir: Down, Delta: 5},
		Command{Dir: Forward, Delta: 8},
		Command{Dir: Up, Delta: 3},
		Command{Dir: Down, Delta: 8},
		Command{Dir: Forward, Delta: 2},
	}
}

func TestPartOne(t *testing.T) {
	commands := testCommands()

	res := partOne(commands)
	expected := 150

	if res != expected {
		t.Fatalf("Want %d Got %d", expected, res)
	}
}

func TestPartTwo(t *testing.T) {
	commands := testCommands()

	res := partTwo(commands)
	expected := 900

	if res != expected {
		t.Fatalf("Want %d Got %d", expected, res)
	}
}
