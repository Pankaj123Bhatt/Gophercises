package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

var totAnswers, correctAnswers, timer int

func inputRecieverFunction(ch1 chan string, quit chan int) {
	for {
		select {
		case <-ch1:
			//fmt.Println(s, "data received !")
		case <-quit:
			fmt.Println("\nTimeout, Quiz finished !")
			fmt.Println("Result :", correctAnswers, "correct out of", totAnswers, "questions !")
			os.Exit(1)
		}
	}
}
func timerFunction(quit chan int) {
	time.Sleep(time.Duration(timer) * time.Second)
	quit <- 3
}
func main() {

	timer = 30

	fmt.Println("Enter the path of the quiz file.")
	path := ""
	fmt.Scanf("%s", &path)
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := csv.NewReader(file)

	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		}
		totAnswers++
	}
	file.Close()

	inp := ""
	fmt.Printf("Welcome to the GOquiz !!! \nDefault time limit is 30 seconds ! Press 1 to change or 0 to continue.")
	fmt.Scanf("%s", &inp)
	if inp == "1" {
		fmt.Printf("Enter new time limit in seconds : ")
		fmt.Scanf("%d", &timer)
		fmt.Println("Timer successfully changed !")
	}
	fmt.Printf("Hit enter to start the quiz !")
	fmt.Scanf("%s", &inp)
	file, err = os.Open(path)

	c1 := make(chan string)
	quit := make(chan int)

	go timerFunction(quit)
	go inputRecieverFunction(c1, quit)

	i := 1
	reader = csv.NewReader(file)
	for {
		val, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println("Q", i, ": What is "+val[0]+" ?")
		fmt.Printf("Your answer -> ")
		fmt.Scanf("%s", &inp)
		c1 <- inp
		if inp == val[1] {
			correctAnswers++
		}
		i++
	}

	fmt.Println("\nQuiz finished !")
	fmt.Println("Result :", correctAnswers, "correct out of", totAnswers, "questions !")
	os.Exit(1)
}
