package main

import "fmt"

func main() {
	users := getNewUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}

	a, b := compare()
	fmt.Println("point1 == point2:", a)
	fmt.Println("point1 == point3:", b)

}

type user struct {
	name string
	age int
	balance float64
	member bool
}

func getNewUsers() []user {
	u1 := user{
		name: "Tracy",
		age: 51,
		balance: 98.43,
		member: true,
	}

	u2 := user{
		age: 19,
		name: "Nick",
	}

	u3 := user{
		"Bob",
		25,
		0,
		false,
	}

	var u4 user

	u4.name = "Sue"
	u4.age = 31
	u4.member = true
	u4.balance = 17.09

	return []user{u1, u2, u3, u4}
}

type point struct {
	x int
	y int
}

func compare() (bool, bool) {
	point1 := struct {
		x int
		y int
	} {
		10,
		10,
	}

	point2 := struct {
		x int
		y int
		// y []int
	}{}
	point2.x = 10
	point2.y = 5
	// point2.y = []int{1,2,3,4,5}

	point3 := point{10, 10}

	return point1 == point2, point1 == point3
}
