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
    // Create a new random generator with a seed
    source := rand.NewSource(time.Now().UnixNano())
    random := rand.New(source)

    // Generate a random number between 1 and 100 using the local random generator
    targetNumber := random.Intn(100) + 1

    // Create a reader for user input
    reader := bufio.NewReader(os.Stdin)

    // Loop until the correct number is guessed
    for {
        fmt.Print("Enter your guess: ")

        // Read the user input
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("An error occurred while reading input. Please try again.")
            continue
        }

        // Trim any newline characters from the input
        input = strings.TrimSpace(input)

        // Convert the input string to an integer
        guess, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input. Please enter an integer.")
            continue
        }

        // Check if the guess is correct
        if guess < targetNumber {
            fmt.Println("Guess something bigger.")
        } else if guess > targetNumber {
            fmt.Println("Guess something smaller.")
        } else {
            fmt.Println("Congratulations! You guessed the correct number.")
            break
        }
    }
}
