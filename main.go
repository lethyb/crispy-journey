package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func setMaxTries() (int, error) {

	const (
		maxTriesEasy    int = 10
		maxTriesMedium  int = 5
		maxTriesHard    int = 3
		maxTriesDefault int = maxTriesMedium
	)

	var difficultyReader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Print("Please select the difficulty level:\n")
	fmt.Printf("[1] Easy (%d chances)\n", maxTriesEasy)
	fmt.Printf("[2] Medium (%d chances)\n", maxTriesMedium)
	fmt.Printf("[3] Hard (%d chances)\n\n", maxTriesHard)

	difficultyInput, err := difficultyReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return maxTriesDefault, err
	}

	difficultyInput = strings.TrimSpace(difficultyInput)

	difficulty, err := strconv.Atoi(difficultyInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number between 1 and 3.")
		return maxTriesDefault, err
	}

	if difficulty < 1 || difficulty > 3 {
		fmt.Println("Invalid choice. Please choose between 1, 2, and 3.")
		return maxTriesDefault, err
	}

	var maxTries int
	switch difficulty {
	case 1:
		maxTries = maxTriesEasy
	case 2:
		maxTries = maxTriesMedium
	case 3:
		maxTries = maxTriesHard
	default:
		maxTries = maxTriesDefault
	}

	return maxTries, nil
}

func generateRandomNumber(min int, max int) int {
	return rand.Intn(max) + min
}

func guessRandomNumber(min int, max int, tries int, maxTries int) (int, error) {

	var choiceNumberReader *bufio.Reader = bufio.NewReader(os.Stdin)

	if tries == 0 {
		fmt.Printf("Enter your guess: ")
	} else {
		fmt.Printf("Try again (%d chances left): ", maxTries-tries)
	}

	choiceNumberInput, err := choiceNumberReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input", err)
		return 0, err
	}

	choiceNumberInput = strings.TrimSpace(choiceNumberInput)

	choiceNumber, err := strconv.Atoi(choiceNumberInput)
	if err != nil {
		fmt.Printf("Invalid choice. Please choose between %d and %d.\n", min, max)
		return 0, err
	}

	if choiceNumber < min || choiceNumber > max {
		fmt.Printf("Invalid choice. Please choose between %d and %d.\n", min, max)
		return 0, err
	}

	return choiceNumber, nil
}

func printIsRandomNumberLessOrGreater(randomNumber int, choiceNumber int) {
	if randomNumber < choiceNumber {
		fmt.Printf("Incorrect! The number is less than %d.\n\n", choiceNumber)
		return
	}

	fmt.Printf("Incorrect! The number is greater than %d.\n\n", choiceNumber)
}

func main() {

	const (
		min int = 1
		max int = 5
	)

	fmt.Printf("Welcome to the Number Guessing Game!\n\n")

	maxTries, err := setMaxTries()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("I'm thinking of a number between %d and %d.\n", min, max)
	fmt.Printf("You have %d chances to guess the correct number.\n\n", maxTries)

	var randomNumber int = generateRandomNumber(min, max)

	var tries int = 0

	for {

		if tries >= maxTries {
			fmt.Printf("You lost.")
			break
		}

		choiceNumber, err := guessRandomNumber(min, max, tries, maxTries)
		if err != nil {
			fmt.Println("Error", err)
			break
		}

		if choiceNumber == randomNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempt(s).", tries+1)
			break
		}

		tries++

		printIsRandomNumberLessOrGreater(randomNumber, choiceNumber)
	}
}
