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
	randomNumber := randomizer.Intn(max-min) + min
	fmt.Println("The random number is", randomNumber)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Number Guessing Game")

	for {
		fmt.Println("Guess a number between 1 and 100")
		fmt.Println("Please input your guess")

		input, _ := reader.ReadString('\n')
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
			fmt.Println("Your guess was correct!")
			break
		}
	}
}
