package main

import (
	"fmt"
	"runtime"
	"unicode"
)

func main()  {
	//if passwordTest("This!I5A") {
	//	fmt.Println("password good")
	//} else {
	//	fmt.Println("password bad")
	//}

	// memoryCheck()
	// floatingNumbers()
	// exploreRune()
	salesTax(0.99, 7.5)
}

func memoryCheck() {
	var list []int8
	//var list []int8
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}

func passwordTest(pw string) bool  {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumber = true
		}

		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasSymbol && hasNumber && hasLower && hasUpper
}

func floatingNumbers() {
	var a int  = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a / 3)
	fmt.Println(b / 3)
	fmt.Println(c / 3)

	fmt.Println("\n")

	fmt.Println((a / 3) * 3)
	fmt.Println((b / 3) * 3)
	fmt.Println((c / 3) * 3)
}

func exploreRune() {
	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}

func salesTax(cost float64, rate float64) {
	tax := cost * (rate / 100)

	fmt.Println("Sales Tat Total: ", tax)
}