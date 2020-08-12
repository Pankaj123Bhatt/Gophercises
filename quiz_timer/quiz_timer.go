package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

var tot, correct, timer int

func f1(ch1 chan string, quit chan int) {
	for {
		select {
		case <-ch1:
			//fmt.Println(s, "data received !")
		case <-quit:
			fmt.Println("\nTimeout, Quiz finished !")
			fmt.Println("Result :", correct, "correct out of", tot, "questions !")
			os.Exit(1)
		}
	}
}
func f2(quit chan int) {
	time.Sleep(time.Duration(timer) * time.Second)
	quit <- 3
}
func main() {

	correct = 0
	timer = 30
	inp := ""
	c1 := make(chan string)
	quit := make(chan int)

	file, err := os.Open("/Users/pankajbhatt/Documents/problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	tot = len(records)
	fmt.Printf("Welcome to the GOquiz !!! \nDefault time limit is 30 seconds ! Press 1 to change or 0 to continue.")
	fmt.Scanf("%s", &inp)
	if inp == "1" {
		fmt.Printf("Enter new time limit in seconds : ")
		fmt.Scanf("%d", &timer)
		fmt.Println("Timer successfully changed !")
	}
	fmt.Printf("Hit enter to start the quiz !")

	fmt.Scanf("%s", &inp)
	go f2(quit)
	go f1(c1, quit)
	for i, val := range records {

		fmt.Println("Q", (i + 1), ": What is "+val[0]+" ?")
		fmt.Printf("Your answer -> ")
		fmt.Scanf("%s", &inp)
		c1 <- inp
		if inp == val[1] {
			correct++
		}
	}
	fmt.Println("\nQuiz finished !")
	fmt.Println("Result :", correct, "correct out of", tot, "questions !")
	os.Exit(1)
}
