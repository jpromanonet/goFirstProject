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
	fmt.Println("Guess a number between 1 and 100")

	secretNumber := generateRandomInteger(1, 100)

	var attempts int
	for {
		attempts++
		fmt.Println("Please input your guess")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		fmt.Println("Your guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again")
		} else {
			fmt.Printf("Correct, you Legend! You guessed right after %d attempts", attempts)
			break
		}
	}
}

func generateRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}