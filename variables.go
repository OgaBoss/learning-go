package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
	medical()
	values()
	getPointers()

	var count int

	add5Value(count)
	count += 4
	add5Pointer(&count)


	a, b := 5, 10
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func medical() {
	first_name := "Damilola"
	family_name := "Adebayo"
	age := 7
	age += 5
	allergy := true

	fmt.Println(first_name)
	fmt.Println(family_name)
	fmt.Println(age)
	fmt.Println(allergy)
}

func values() {
	var count int
	var discount float64
	var emails []string
	var startTime time.Time
	fmt.Printf("Count : %#v \n", count)
	fmt.Printf("Count : %#v \n", discount)
	fmt.Printf("Count : %#v \n", emails)
	fmt.Printf("Count : %#v \n", startTime)
}

func getPointers() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time  : %#v\n", t)

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}

	if t != nil {
		fmt.Printf("time : %#v\n", *t)
		fmt.Printf("time : %#v\n", t.String())
	}
}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value :", count)
}

func add5Pointer(count *int) {
	*count += 5
	fmt.Println("add5Pointer :", *count)
}