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

// Represents a single leaderboard record
type LeaderboardEntry struct {
	Difficulty string `json:"difficulty"`
	Attempts   int    `json:"attempts"`
	Data       string `json:"date"`
}

// Store records for each difficulty
type Leaderboard struct {
	Easy    []LeaderboardEntry `json:"easy"`
	Medium  []LeaderboardEntry `json:"medium"`
	Hard    []LeaderboardEntry `json:"hard"`
	MaxSize int                `json:"max_size"` // Top N entries
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Number Guessing Game")

	for {
		handleGame(reader)

		fmt.Print("\nWould you like to play again? (y/n): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer != "y" && answer != "yes" {
			fmt.Println("Thanks for playing")
			break
		}
	}
}

func handleGame(reader *bufio.Reader) {
	min, max := 1, 100
	randomNumber := rand.Intn(max-min+1) + min

	maxAttempts := handleDifficulty(reader)

	fmt.Printf("\nGuess a number between %d and %d\n", min, max)
	fmt.Printf("You have %d attempts.\n\n", maxAttempts)

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Attempt %d/%d - Enter your guess: ", attempts, maxAttempts)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			attempts--
			continue
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			attempts--
			continue
		}

		if guess > randomNumber {
			fmt.Println("Your guess is bigger than the random number. Try again")
		} else if guess < randomNumber {
			fmt.Println("Your guess is smaller than the random number. Try again")
		} else {
			fmt.Printf("\nCongratulations! You guessed the number in %d attempts!\n", attempts)
			break
		}

		if attempts == maxAttempts {
			fmt.Println("\nGame over")
			fmt.Printf("The correct number was %d\n", randomNumber)
		}
	}
}

func handleDifficulty(reader *bufio.Reader) int {
	for {
		fmt.Println("\nChoose difficulty:")
		fmt.Println("1. Easy   (10 attempts)")
		fmt.Println("1. Medium (5 attempts)")
		fmt.Println("1. Hard   (3 attempts)")
		fmt.Print("Enter choice:")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1", "easy":
			return 10
		case "2", "medium":
			return 5
		case "3", "hard":
			return 3
		default:
			fmt.Println("Invalid choise. Please try again.")
		}
	}
}

func handleLeaderboard() *Leaderboard {
	return &Leaderboard{
		Easy:    []LeaderboardEntry{},
		Medium:  []LeaderboardEntry{},
		Hard:    []LeaderboardEntry{},
		MaxSize: 10, // Keep top 10 scores per difficulty
	}
}
