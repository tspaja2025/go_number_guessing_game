package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min, max := 1, 100
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	randomNumber := randomizer.Intn(max-min) + min
	fmt.Println("The random number is", randomNumber)

	fmt.Println("Welcome to Number Guessing Game")
	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Please input your guess")

	var guess int

	fmt.Scan(&guess)

	if guess > randomNumber {
		fmt.Println("Your guess is bigger than the random number. Try again")
	} else if guess < randomNumber {
		fmt.Println("Your guess is smaller than the random number. Try again")
	} else {
		fmt.Println("Your guess was correct!")
	}
}
