package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Enter the path of the quiz file.")
	path := ""
	fmt.Scanf("%s", &path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)

	fmt.Printf("Hit enter to start !")
	input := ""
	fmt.Scanf("%s", &input)
	totalQuestions := 0
	correctAnswers := 0
	i := 1
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println("Q", i, ": What is "+records[0]+" ?")
		fmt.Printf("Your answer -> ")
		fmt.Scanf("%s", &input)
		if input == records[1] {
			correctAnswers++
		}
		totalQuestions++
		i++
	}

	fmt.Println("Quiz finished !")
	fmt.Println("Result :", correctAnswers, "correct out of", totalQuestions, "questions !")
}
