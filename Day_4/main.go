package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("Guess the number between 1 and 100")
	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)
	var chances int
	switch choice {
	case 1:
		fmt.Println("Easy")
		chances = 10
	case 2:
		fmt.Println("Medium")
		chances = 7
	case 3:
		fmt.Println("Hard")
		chances = 5
	default:
		fmt.Println("Invalid choice. Defaulting to Medium difficulty (7 chances).")
		chances = 7 // Default to medium if invalid choice
	}
	fmt.Printf("You have selected %d chances.\n", chances) // Corrected message
	var randomnumber int
	rand.Seed(time.Now().UnixNano())
	randomnumber = rand.Intn(100) + 1
	// fmt.Printf("the random number is %d\n", randomnumber) // Hint for debugging, turned off for the real game!

	for i := 0; i < chances; i++ {
		fmt.Printf("\nYou have %d chances left.\n", chances-i)
		fmt.Println("Enter your guess: ")
		var guess int
		fmt.Scan(&guess)

		if guess == randomnumber {
			fmt.Println("You win! 🏆")
			break
		} else if guess < randomnumber {
			fmt.Println("Too low! 📉")
		} else {
			fmt.Println("Too high! 📈")
		}

		if i == chances-1 {
			fmt.Printf("\nGame Over! The number was %d.\n", randomnumber)
		}
	}
	// The final random number reveal is now handled inside the loop for game over condition,
	// or implicitly by the win condition. This line can be removed or kept for explicit reveal.
	// fmt.Printf("the random number is %d\n", randomnumber)

}
