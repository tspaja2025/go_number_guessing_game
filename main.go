package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	min, max := 1, 100

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := randomizer.Intn(max-min+1) + min
	fmt.Println("The random number is", randomNumber)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Number Guessing Game")
	fmt.Println("Guess a number between 1 and 100")

	for attemps := 1; ; attemps++ {
		fmt.Printf("Attempt %d: Please enter your guess: ", attemps)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		if guess > randomNumber {
			fmt.Println("Your guess is bigger than the random number. Try again")
		} else if guess < randomNumber {
			fmt.Println("Your guess is smaller than the random number. Try again")
		} else {
			fmt.Printf("Congratulations you guessed the number! It took %d attempts for you to guess the number.", attemps)
			break
		}

		if attemps == 3 {
			fmt.Println("Game over")
			fmt.Printf("The correct number is %d\n", randomNumber)
			break
		}
	}
}
