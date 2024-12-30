// Case Study:
// Develop an online examination system where users can take a quiz.
// 1. Question Bank: Define a slice of structs to store questions. Each question
// should have a question string, options (array), and the correct answer.
// 2. Take Quiz: Use loops to iterate over questions and display them one by one.
// Allow the user to select an answer by entering the option number.
// o Use continue to skip invalid inputs and prompt the user again.
// o Use break to exit the quiz early if the user enters a specific command
// (e.g., "exit").
// 3. Score Calculation: After the quiz, calculate the user's score and display it. Use
// conditions to classify performance (e.g., "Excellent", "Good", "Needs
// Improvement").
// 4. Error Handling: Handle errors like invalid input during the quiz (e.g., entering a
// non-integer value for an option).

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

var questionBank = []Question{
	{
		Question: "What is the capital of France?",
		Options:  [4]string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"},
		Answer:   3,
	},
	{
		Question: "Which is the largest planet in our solar system?",
		Options:  [4]string{"1. Earth", "2. Jupiter", "3. Mars", "4. Venus"},
		Answer:   2,
	},
	{
		Question: "What is 2 + 2?",
		Options:  [4]string{"1. 3", "2. 4", "3. 5", "4. 6"},
		Answer:   2,
	},
}

func takeQuiz() int {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Quiz! Type 'exit' anytime to quit.")

	for i, q := range questionBank {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		for {
			fmt.Print("Enter your answer (1-4) or 'exit': ")
			if !scanner.Scan() {
				fmt.Println("Error reading input.")
				continue
			}
			input := scanner.Text()

			if strings.ToLower(input) == "exit" {
				fmt.Println("Exiting the quiz...")
				return score
			}

			answer, err := strconv.Atoi(input)
			if err != nil || answer < 1 || answer > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}

			if answer == q.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong answer.")
			}
			break
		}
	}

	return score
}

func classifyPerformance(score, total int) string {
	switch {
	case score == total:
		return "Excellent"
	case score >= total/2:
		return "Good"
	default:
		return "Needs Improvement"
	}
}

func main() {
	totalQuestions := len(questionBank)
	score := takeQuiz()

	fmt.Printf("\nYou scored %d out of %d.\n", score, totalQuestions)
	performance := classifyPerformance(score, totalQuestions)
	fmt.Printf("Your performance: %s\n", performance)
}
