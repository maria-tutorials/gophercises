# gophercises-quiz-game

Small cli program that reads a quiz from a CSV file and then asks the questions to the user, whilst keeping track of responses.

`./quiz --help`

### csv
the first column is a question and the second column in the same row is the answer to that question.

## Part 1

Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

## Part 2

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded.

Users should be asked to press enter (or some other key) before the timer starts.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

## Extras
1. Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect

2. Add an option (a new flag) to shuffle the quiz order each time it is run.
