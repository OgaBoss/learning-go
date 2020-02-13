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
	control()
	fmt.Println(deleteFromSlice())

	// Map
	fmt.Println("Users:", getUsers())
	printUser()



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
func control() {
	l1, l2, l3 := linked()
	fmt.Println("Linked :", l1, l2, l3)
	nl1, nl2 := noLink()
	fmt.Println("No Link :", nl1, nl2)
	cl1, cl2 := capLinked()
	fmt.Println("Cap Link :", cl1, cl2)
	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)
	copy1, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copy1, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)
	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}

//In the Linked scenario, we made a simple copy of the first slice and then a
//range copy of it. While the slices themselves are distinct and are no longer the same
//slices, in reality, it doesn't make a difference to the data they hold. Each of the slices
//pointed to the same hidden array, so when we made a change to the first slice, it
//affected all of the slices.
func linked() (int, int, int) {
	s1 := []int{1,2,3,4,5}
	s2 := s1
	s3 := s1[:]
	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

//In the No Link scenario, the setup was the same for the first and second slice, but before
//we made a change to the first slice, we appended a value to it. When we appended this
//value to it, in the background, Go needed to create a new array to hold the now large
//number of values. Since we were appending to the first slice, its pointer was to look at
//the new, bigger slice. The second slice doesn't get its pointer updates. That's why, when
//the first slice had its value change, the second slice wasn't affected. The second slice
//isn't pointing to the same hidden array anymore, meaning they are not linked.
func noLink() (int, int)  {
	s1 := []int{1,2,3,4,5}
	s2 := s1

	s1 = append(s1, 6)

	s1[3] = 99

	return s1[3], s2[3]
}

//For the Cap Link scenario, the first slice was defined using make and with an oversized
//capacity. This extra capacity meant that when the first slice had a value appended to it,
//there was already extra room in the hidden array. This extra capacity means there's no
//need to replace the hidden array. The effect was that when we updated the value on the
//first slice, it and the second slice were still pointing to the same hidden array, meaning
//the change affects both.
func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1
	s1 = append(s1, 6)

	s1[3] = 99

	return s1[3], s2[3]
}

//In the Cap No Link scenario, the setup was the same as the previous scenario, but when
//we appended values, we appended more values than there was available capacity. Even
//though there was extra capacity, there was not enough, and the hidden array in the first
//slice got replaced. The result was that the link between the two slices broke.
func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1
	s1 = append(s1, []int{10: 11}...)

	s1[3] = 99

	return s1[3],s2[3]
}

//In Copy No Link, we used the built-in copy function to copy the value for us. While this
//does copy the values into a new hidden array, copy won't change the length of the slice.
//This fact means that the destination slice must be the correct length before you do the
//copy. You don't see copy much in real-world code; this could be because it's easy to
//misuse it.
func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))

	copied := copy(s2, s1)

	s1[3] = 99

	return s1[3], s2[3], copied
}

//Lastly, with Append No Link, we use append to do something similar to copy but without
//having to worry about the length. This method is the most commonly seen in realworld
//code when you need to ensure you get a copy of the values that are not linked to
//the source. This is easy to understand since append gets used a lot and it's a one-line
//solution. There is one slightly more efficient solution that avoids the extra memory
//allocation of the empty slice in the first argument of append.
func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)

	// OR
	// This uses the seldom-used slice range notation of
	// <slice>[<low>:<high>:<capacity>]. With the current Go compiler, this is the most
	// memory-efficient way to copy a slice.
	// s2 := append(s1[:0:0], s1...)

	s1[3] = 99

	return s1[3], s2[3]
}

func deleteFromSlice() []string {
	arr := []string{"Good", "Good", "Bad", "Good", "Good"}

	s2 := arr[:2]

	return append(s2, arr[3:]...)
}

// Maps
func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
	}

	users["073"] = "Tracy"

	return users
}

func getUser(id string) (string, bool) {
	users := getUsers()

	user, exists := users[id]

	return user, exists
}

func printUser(){
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	name, exists := getUser(userID)

	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userID)
		for key, value := range getUsers() {
			fmt.Println(" ID:", key, "Name:", value)
		}
		os.Exit(1)
	}

	fmt.Println("Name:", name)
}