package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//General Structure of Problems
type Problem struct {
	question string
	answer   string
}

func main() {
	csvFile := flag.String("file", "problems.csv", "CSV file in the format of 'question,answer' ")
	timeLimit := flag.Int("time", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Not able to open the file: %s", *csvFile))
	}

	parsedFile := csv.NewReader(file)
	lines, err := parsedFile.ReadAll()
	if err != nil {
		exit("Not able to read the data from the file.")
	}
	quizProblems := parseQuestionsFromCSV(lines)
	// fmt.Println(quizProblems)
	fmt.Printf("Starting quiz with %d questions, time limit: %d\n", len(quizProblems), *timeLimit)
	startQuiz(quizProblems, *timeLimit)
}

func parseQuestionsFromCSV(lines [][]string) []Problem {
	quizProblems := make([]Problem, len(lines))
	for index, currentLine := range lines {
		quizProblems[index] = Problem{
			question: currentLine[0],
			answer:   strings.TrimSpace(currentLine[1]),
		}
	}
	return quizProblems
}

func startQuiz(quizProblems []Problem, timeLimit int) {
	//Create a timer for the quiz
	//After the specified time is elapsed, timer will add data to the channel.
	quizTimer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	score := 0
	for index, problem := range quizProblems {
		fmt.Printf("Question: #%d:\t%s\t", index+1, problem.question)
		//Create a channel to receive answers
		//This is create a seperate channel for collecting answers from users.
		//This needs to be inside the loop since we need to get answers for every question
		//Scanf will wait till some text is entered by the user. Once the text is entered and 'Enter' is pressed, the value will be inserted into the channel.
		answerChannel := make(chan string)
		go func() {
			var currentAnswer string
			fmt.Scanf("%s\n", &currentAnswer)
			answerChannel <- currentAnswer
		}()
		select {
		case <-quizTimer.C:
			fmt.Printf("\nTime out, Quiz Complete. You Scored %d out of %d", score, len(quizProblems))
			return
		case currentAnswer := <-answerChannel:
			score += validateAnswer(currentAnswer, problem.answer)
		}
	}
	fmt.Printf("\nQuiz Complete. You Scored %d out of %d", score, len(quizProblems))
}

func validateAnswer(currentAnswer string, correctAnswer string) int {
	if currentAnswer == correctAnswer {
		return 1
	} else {
		return 0
	}
}
func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
