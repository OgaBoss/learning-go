package main

import (
	"fmt"
	"time"
)

func main() {
	input := 5

	if input%2 == 0 {
		fmt.Println(input, "is even")
	}

	if input%2 == 1 {
		fmt.Println(input, "is odd")
	}

	//caseSwitch()
	//fizzBuzz()
	//mapLoop()
	maxCount()
}

func caseSwitch() {
	dayBorn := time.Monday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on a weekend")
	default:
		fmt.Println("Error, day born not valid")
	}
}

func fizzBuzz() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i,"FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func mapLoop() {
	config := map[string]string{
		"debug": "1",
		"logLevel": "warn",
		"version": "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}

func maxCount() {
	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up": 4,
	}

	maxWord := ""
	maxCount := 0

	for key, value := range words {
		if value > maxCount {
			maxWord = key
			maxCount = value
		}
	}

	fmt.Println("Most popular word:", maxWord)
	fmt.Println("With count of", maxCount)
}

func bubbleSort() {

}