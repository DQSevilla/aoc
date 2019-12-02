package main

import (
	"fmt"
	"io"
	"os"
)

const inputFilePath = "input"

// TotalFuelRequirements computes the sum of all fuel requirements for all
// modules based on their mass. Solves part one.
func TotalFuelRequirements(module_masses []int) (total int) {
	for _, mass := range module_masses {
		total += mass/3 - 2
	}
	return
}

// main driver for AOC 2019 #1
func main() {
	var module_masses []int

	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		var mass int
		_, err = fmt.Fscanln(file, &mass)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		module_masses = append(module_masses, mass)
	}

	fmt.Println("Total fuel requirements for all modules on spacecraft:",
		TotalFuelRequirements(module_masses))
}
