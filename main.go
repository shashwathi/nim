package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Clearly, the human could have played better in the above game.
	// The computer randomly removes 1 or 2 sticks, but cannot remove more sticks than are left.
	// The human is prompted at each turn for how many sticks he or she wants to remove.
	// Be careful! A human might enter 5 if 5 sticks are left, and if you are not careful, the human could win by playing in that way.
	// Don't accept the user's input if it is illegal. Continue prompting until you get a valid input.
	// Start your work by creating a Nim class in the lab2 package of the labs source folder in your repository.
	// Use ArgsProcessor to prompt for inputs.
	// Your program must continue play until somebody (computer or human) wins.
	// When you demo this assignment, be prepared to discuss how you would implement a dominant strategy with your TA

	//On each turn, a player must remove either 1 or 2 sticks from the pile. The goal of the game is to be the player who removes the last stick.
	// desired output
	// Computer starts
	// Round 0, 7 sticks at start, computer took 2, so 5 sticks remain
	// Round 1, 5 sticks at start, human took 1, so 4 sticks remain
	// Round 2, 4 sticks at start, computer took 2, so 2 sticks remain
	// Round 3, 2 sticks at start, human took 1, so 1 sticks remain
	// Round 4, 1 sticks at start, computer took 1, so 0 sticks remain
	// Computer wins

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter minimum heap size: ")
	// Read from stdin
	heapSize, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading heap size: %v", err)
		return
	}
	// Trim new line from input
	heapSize = strings.TrimSuffix(heapSize, "\n")
	// Convert the heao size to int
	s, err := strconv.ParseInt(heapSize, 10, 64)
	if err != nil {
		fmt.Printf("Error reading heap size: %v", err)
		return
	}
	// convert heap size into int from int64
	heapS := int(s)
	var i = 0
	for i = 0; i < heapS; i++ {
		// use random generator for computer to pick the number
		r := rand.Intn(heapS)
		// update the heap size by removing the computer selected sticks
		updatedHeapSize := heapS - r
		fmt.Printf("\n Round %d, %d sticks at start, computer took %d, so %d sticks remains\n", i, heapS, r, (heapS - r))
		// update the heap size
		heapS = updatedHeapSize
		// Player will choose the sticks
		i++
		playerChoice, err := readInput()
		if err != nil {
			fmt.Printf("\nError reading input: %v", err)
			return
		}

		if playerChoice < heapS {
			updatedHeapSize = heapS - playerChoice
			fmt.Printf("\n Round %d, %d sticks at start, player took %d, so %d sticks remains\n", i, heapS, playerChoice, (heapS - playerChoice))
			// update the heap size
			heapS = updatedHeapSize
		} else {
			fmt.Printf("Error in player choice. Retry again")
			return
		}
	}
	// USe heap size to decide if computer / human won tge game.
	if heapS >= 1 {
		fmt.Printf("\n Round %d, %d sticks at start, computer took %d, so %d sticks remains\n", i, heapS, heapS, 0)
		fmt.Println("\nComputer wins")
	} else {
		fmt.Println("\nHuman wins")
	}
	return
}

func readInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nPlayer pick sticks: ")
	// Read from stdin
	heapSize, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading heap size: %v", err)
		return 0, err
	}
	// Trim new line from input
	heapSize = strings.TrimSuffix(heapSize, "\n")
	// Convert the heao size to int
	s, err := strconv.ParseInt(heapSize, 10, 64)
	if err != nil {
		fmt.Printf("Error reading heap size: %v", err)
		return 0, nil
	}
	return int(s), nil
}
