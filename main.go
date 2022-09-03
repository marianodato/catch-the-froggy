package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numberOfBoxes    = 100
	firstBoxPosition = 1
)

func main() {
	rand.Seed(time.Now().UnixNano())
	maximumSolutionSteps := getMaximumSolutionSteps()
	fmt.Printf("Maximum solution steps: %v\n", maximumSolutionSteps)
	for i := 0; i < 1000000; i++ {
		result := testSolution(maximumSolutionSteps)
		if !result {
			fmt.Print("Solution did not work!")
			return
		}
	}
	fmt.Print("Solution worked!")
	return
}

// Generates the sequence {n-1, n-1, n-2, n-3, ..., 2, 2, 3, 4, ..., n} with n = numberOfBoxes
// If numberOfBoxes == 1 the solution sequence is {1}
func getMaximumSolutionSteps() []int {
	if numberOfBoxes == 1 {
		return []int{firstBoxPosition}
	}
	const solutionSize = (numberOfBoxes - 1) * 2
	maximumSolutionSteps := make([]int, solutionSize)
	for i := range maximumSolutionSteps {
		if i == 0 {
			maximumSolutionSteps[i] = numberOfBoxes - 1
		} else if i <= numberOfBoxes-2 {
			maximumSolutionSteps[i] = numberOfBoxes - i
		} else if i == numberOfBoxes-1 {
			maximumSolutionSteps[i] = maximumSolutionSteps[i-1]
		} else {
			maximumSolutionSteps[i] = maximumSolutionSteps[i-1] + 1
		}
	}
	return maximumSolutionSteps
}

func testSolution(maximumSolutionSteps []int) bool {
	froggiePosition := getRandomPosition(firstBoxPosition, numberOfBoxes)
	// fmt.Printf("Initial froggie position: %d\n", froggiePosition)
	for i := range maximumSolutionSteps {
		chosenPosition := maximumSolutionSteps[i]
		// fmt.Printf("Chosen position: %d\n", chosenPosition)
		if chosenPosition == froggiePosition {
			// fmt.Println("Froggie caught!")
			return true
		}
		froggiePosition = getNextFroggiePosition(froggiePosition)
		// fmt.Printf("Froggie position: %d\n", froggiePosition)
	}
	// fmt.Println("Could not catch froggie!")
	return false
}

func getRandomPosition(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func getNextFroggiePosition(froggiePosition int) int {
	const leftDirection = -1
	const rightDirection = 1
	nextDirection := []int{leftDirection, rightDirection}
	chosenIndex := rand.Intn(len(nextDirection))
	nextFroggiePosition := froggiePosition + nextDirection[chosenIndex]
	if nextFroggiePosition < firstBoxPosition || nextFroggiePosition > numberOfBoxes {
		const invertDirection = -1
		nextFroggiePosition = froggiePosition + nextDirection[chosenIndex]*(invertDirection)
	}
	return nextFroggiePosition
}
