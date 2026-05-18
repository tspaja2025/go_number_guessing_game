package main

import (
	"bufio"
	"encoding/json"
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
	Date       string `json:"date"`
}

// Store records for each difficulty
type Leaderboard struct {
	Easy    []LeaderboardEntry `json:"easy"`
	Medium  []LeaderboardEntry `json:"medium"`
	Hard    []LeaderboardEntry `json:"hard"`
	MaxSize int                `json:"max_size"` // Top N entries
}

const leaderboardFile = "leaderboard.json"

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
			fmt.Println("Thanks for playing!")
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

func handleGetDifficultyName(difficultyCode int) string {
	switch difficultyCode {
	case 10:
		return "easy"
	case 5:
		return "medium"
	case 3:
		return "hard"
	default:
		return "unknown"
	}
}

func handleSortEntries(entries []LeaderboardEntry) {
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].Attempts > entries[j].Attempts {
				entries[i], entries[j] = entries[j], entries[i]
			}
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

func handleLoadLeaderboard(filename string) (*Leaderboard, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return handleLeaderboard(), nil
		}
		return nil, err
	}
	defer file.Close()

	var leaderboard Leaderboard
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&leaderboard)
	if err != nil {
		return handleLeaderboard(), nil
	}

	if leaderboard.MaxSize == 0 {
		leaderboard.MaxSize = 10
	}

	return &leaderboard, nil
}

func (leaderboard *Leaderboard) SaveLeaderboard(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(leaderboard)
}

func (leaderboard *Leaderboard) AddEntry(difficulty string, attempts int) bool {
	entry := LeaderboardEntry{
		Difficulty: difficulty,
		Attempts:   attempts,
		Date:       time.Now().Format("2006-01-02 15:04:05"),
	}

	var entries *[]LeaderboardEntry
	switch strings.ToLower(difficulty) {
	case "easy":
		entries = &leaderboard.Easy
	case "medium":
		entries = &leaderboard.Medium
	case "hard":
		entries = &leaderboard.Hard
	default:
		return false
	}

	currentBest := leaderboard.GetBestAttempt(difficulty)
	if currentBest == 0 || attempts < currentBest {
		*entries = append(*entries, entry)

		handleSortEntries(*entries)

		if len(*entries) > leaderboard.MaxSize {
			*entries = (*entries)[:leaderboard.MaxSize]
		}

		return true
	}

	return false
}

func (leaderboard *Leaderboard) GetBestAttempt(difficulty string) int {
	var entries []LeaderboardEntry
	switch strings.ToLower(difficulty) {
	case "easy":
		entries = leaderboard.Easy
	case "medium":
		entries = leaderboard.Medium
	case "hard":
		entries = leaderboard.Hard
	default:
		return 0
	}

	if len(entries) == 0 {
		return 0
	}

	return entries[0].Attempts
}

func (leaderboard *Leaderboard) LeaderboardDisplay() {
	fmt.Println("\n", strings.Repeat("=", 50))
	fmt.Println("LEADERBOARD")
	fmt.Println(strings.Repeat("=", 50))

	displayDifficulty := func(name string, entries []LeaderboardEntry) {
		if len(entries) == 0 {
			fmt.Printf("\n%s: No records yet\n", strings.Title(name))
			return
		}

		fmt.Printf("\n %s Difficulty (Best: %d attempts)\n", strings.Title(name), entries[0].Attempts)
		fmt.Println(strings.Repeat("-", 40))

		for i, entry := range entries {
			medal := ""
			switch i {
			case 0:
				medal = "🥇 "
			case 1:
				medal = "🥈 "
			case 2:
				medal = "🥉 "
			default:
				medal = "  "
			}
			fmt.Printf("%s#%d: %d attempts (%s)\n", medal, i+1, entry.Attempts, entry.Date)
		}
	}

	displayDifficulty("easy", leaderboard.Easy)
	displayDifficulty("medium", leaderboard.Medium)
	displayDifficulty("hard", leaderboard.Hard)

	fmt.Println(strings.Repeat("=", 50))
}
