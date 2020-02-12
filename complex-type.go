package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", defineArray())

	// Arrays
	comp1, comp2 := compArrays()
	fmt.Println("[5]int == [5]int{0} :", comp1)
	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0}:", comp2)

	var arr [10]int
	arr = fill(arr)
	fmt.Println(arr)

	// Slices
	sliceMain()
}


// Arrays
func defineArray() [10]string {
	var arr [10]string
	return arr
}

func compArrays() (bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}

	return arr1 == arr2, arr1 == arr3
}

func fill(array [10]int) [10]int {
	for i := 0; i < len(array); i++ {
		array[i] = i+1
	}

	return array
}

//Slices
func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}

	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func findLongestString(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}

	return longest
}

func getLocals(extraLocales []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocales...)

	return locales
}

func sliceMain() {
	if longest := findLongestString(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}

	locales := getLocals(getPassedArgs(3))
	fmt.Println("Locales to use:", locales)
}

// Controlling internal slice behaviour
func linked() (int, int, int) {
	s1 := []int{1,2,3,4,5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

func noLink() (int, int)  {
	s1 := []int{1,2,3,4,5}
	s2 := s1

	s1 = append(s1, 6)

	s1[3] = 99

	return s1[3], s2[3]
}