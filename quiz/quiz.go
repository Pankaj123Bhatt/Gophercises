package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/pankajbhatt/Documents/problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Printf("Hit enter to start !")
	inp := ""
	fmt.Scanf("%s", &inp)
	tot := 0
	correct := 0

	for i, val := range records {
		fmt.Println("Q", (i + 1), ": What is "+val[0]+" ?")
		fmt.Printf("Your answer -> ")
		fmt.Scanf("%s", &inp)
		if inp == val[1] {
			correct++
		}
		tot++
	}
	fmt.Println("Quiz finished !")
	fmt.Println("Result :", correct, "correct out of", tot, "questions !")

}
