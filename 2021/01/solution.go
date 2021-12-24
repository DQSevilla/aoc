package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne(depths []int) int {
	var numIncreases int

	prev := depths[0]
	for _, depth := range depths[1:] {
		if depth > prev {
			numIncreases++
		}

		prev = depth
	}

	return numIncreases
}

func partTwo(depths []int) int {
	var numIncreases int

	for i, prev := range depths[:len(depths)-3] {
		depth := depths[i+3]
		if depth > prev {
			numIncreases++
		}
	}

	return numIncreases
}

func readInput(filename string) ([]int, error) {
	var nums []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nums, err
		}
		nums = append(nums, num)
	}

	return nums, scanner.Err()
}

func main() {
	depths, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	solutionOne := partOne(depths)
	//depths = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	solutionTwo := partTwo(depths)
	fmt.Printf("Solution One: %d", solutionOne)
	fmt.Println()
	fmt.Printf("Solution Two: %d", solutionTwo)
}
