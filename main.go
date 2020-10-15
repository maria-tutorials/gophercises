package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

type scoreboard struct {
	total   int64
	correct int64
}

var us scoreboard

var quiz []problem

func main() {
	fn := flag.String("csv", "problems.csv", "the quiz file as a csv")
	//time := flag.Int("timeout", 30, "total time to finish the quiz, in seconds")
	flag.Parse()

	lines := readFile(fn)
	buildQuiz(lines)

	for i, q := range quiz {
		fmt.Printf("Question #%d: \n", (i + 1))

		handleQuestion(q)
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

	return lines
}

func buildQuiz(l [][]string) {
	us.total = int64(len(l))
	for _, line := range l {
		p := problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
		quiz = append(quiz, p)
	}
}

func handleQuestion(p problem) {
	fmt.Printf("%s = ?\n", p.question)

	answer := ""

	fmt.Scanf("%s\n", &answer)

	if answer == p.answer {
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

func exit(msg string, err error) {
	log.Fatalln(msg, "\n\twith error:", err)
}
