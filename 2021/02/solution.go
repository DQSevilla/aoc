package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	// Direction describes which way the submarine travels underwater
	Direction int
	// Command describes the submarine's direction and positional change
	Command struct {
		Dir   Direction
		Delta int
	}
)

const (
	// Forward direction
	Forward Direction = 0
	// Down direction
	Down Direction = 1
	// Up direction
	Up Direction = 2
)

func (c Command) String() string {
	return fmt.Sprintf("Dir: %d, Delta: %d", c.Dir, c.Delta)
}

func partOne(commands []Command) int {
	var hp, depth int

	for _, command := range commands {
		switch command.Dir {
		case Forward:
			hp += command.Delta
		case Down:
			depth += command.Delta
		default:
			depth -= command.Delta
		}
	}

	return hp * depth
}

func partTwo(commands []Command) int {
	var aim, hp, depth int

	for _, command := range commands {
		switch command.Dir {
		case Forward:
			hp += command.Delta
			depth += aim * command.Delta
		case Down:
			aim += command.Delta
		case Up:
			aim -= command.Delta
		}
	}

	return hp * depth
}

// StringToDirection maps the raw direction text to the enum constants
func StringToDirection(dir string) (Direction, bool) {
	val, ok := map[string]Direction{
		"forward": Forward,
		"down":    Down,
		"up":      Up,
	}[dir]
	return val, ok
}

func readInput(filename string) ([]Command, error) {
	var commands []Command

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir, ok := StringToDirection(parts[0])
		if !ok {
			return nil, fmt.Errorf("Invalid direction: %s", parts[0])
		}

		delta, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		commands = append(commands, Command{Dir: dir, Delta: delta})
	}

	return commands, scanner.Err()
}

func main() {
	commands, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solutionOne := partOne(commands)
	solutionTwo := partTwo(commands)

	fmt.Printf("Solution One: %d\n", solutionOne)
	fmt.Printf("Solution Two: %d\n", solutionTwo)
}
