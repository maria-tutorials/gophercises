package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const DEFAULT_FILENAME = "problems.csv"
const DEFAULT_LIMIT = 30

type problem struct {
	question, answer string
}

type scoreboard struct {
	total, correct int64
}

var us scoreboard
var quiz []problem

func main() {
	fn := flag.String("csv", DEFAULT_FILENAME, "the quiz file as a csv")
	limit := flag.Int("limit", DEFAULT_LIMIT, "total time to finish the quiz, in seconds")
	shuffle := flag.Bool("shuffle", false, "randomize the questions or present them in order")
	flag.Parse()

	lines := readFile(fn)
	buildQuiz(lines)

	if *shuffle {
		shuffleQuiz()
	}

	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	answerChan := make(chan string)

	for i, p := range quiz {
		handleQuestion(i, p, answerChan)

		select { // lets a goroutine wait on communication operations.
		case <-timer.C:
			fmt.Println("\nTime's up")
			printScore()
			return
		case answer := <-answerChan:
			handleAnswer(answer, p.answer)
		}

	}

	printScore()
}

func readFile(fn *string) [][]string {
	file, err := os.Open(*fn)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s", *fn), err)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll() //Read and parse one by one?
	if err != nil {
		exit(fmt.Sprintf("Failed to read the file: %s", *fn), err)
	}

	if len(lines) == 0 {
		exit(fmt.Sprintf("There are no questions the file: %s", *fn), nil)
	}

	return lines
}

func buildQuiz(l [][]string) {
	us.total = int64(len(l))
	for _, line := range l {
		p := problem{
			question: line[0],
			answer:   cleanStrings(line[1]),
		}
		quiz = append(quiz, p)
	}
}

func shuffleQuiz() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := range quiz {
		newpos := r.Intn(len(quiz) - 1)

		quiz[i], quiz[newpos] = quiz[newpos], quiz[i]
	}
}

func handleQuestion(i int, p problem, answerChan chan string) {
	fmt.Printf("Question #%d: \n", (i + 1))
	fmt.Printf("%s = ?\n", p.question)
	go func() {
		answer := ""
		fmt.Scanf("%s\n", &answer)
		answerChan <- answer
	}()
}

func handleAnswer(userAnswer string, correctAnswer string) {
	if cleanStrings(userAnswer) == correctAnswer {
		us.correct++
		fmt.Println("CORRECT!")
	} else {
		fmt.Println("Sorry that was wrong!")
	}
}

func printScore() {
	score := (us.correct / us.total) * 100

	fmt.Printf("\n\tYou answered %d out of %d correctly.", us.correct, us.total)
	fmt.Printf("\n\tThat's %d%%. ", score)

	switch {
	case score <= 10:
		fmt.Println("Let's try harder next time")
	case score > 10 && score < 50:
		fmt.Println("Better luck next time")
	case score > 50 && score < 90:
		fmt.Println("Good job, that's a passing grade")
	case score > 90 && score < 100:
		fmt.Println("Well done!")
	case score == 100:
		fmt.Println("Fantastic job :D")
	}

}

func cleanStrings(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func exit(msg string, err error) {
	log.Fatalln(msg, "\n\twith error:", err)
}
