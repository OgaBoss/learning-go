package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	users := getNewUsers()
	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}

	a, b := compare()
	fmt.Println("point1 == point2:", a)
	fmt.Println("point1 == point3:", b)

	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}

	fmt.Print(convert())

	caller()
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

// Struct Composition Using Embedding
type name string

type location struct {
	x int
	y int
}

type size struct {
	width int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot  {
	var dot1 dot

	dot2 := dot{}
	dot2.name = "A"
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  34,
			height: 58,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1, dot2, dot3, dot4}
}

// Numeric Type Conversion
func convert() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	// convert from smaller int to larger int
	m := fmt.Sprintf("int8 = %v > in64 = %v\n", i8, int64(i8))

	// Now, we'll convert from an int that's 1 above int8's maximum size. This will cause
	// an overflow to int8's minimum size:
	m += fmt.Sprintf("int = %v > in8 = %v\n", i, int8(i))

	// Next, we'll convert out int8 into a float64. This doesn't cause an overflow and the
	// data is unchanged:
	m += fmt.Sprintf("int8 = %v > float32 = %v\n", i8, float64(i8))

	// Here, we'll convert a float into an int. All the decimal data is lost but the whole
	// number is kept as is:
	m += fmt.Sprintf("float64 = %v > int = %v\n", f64, int(f64))

	return m
}


// Type Assertion
func doubler(v interface{}) (string, error) {
	if i, ok := v.(int); ok {
		return fmt.Sprint(i*2), nil
	}

	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("unsupported type passed")
}

func caller() {
	res, _ := doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	_, err := doubler(true)
	fmt.Println("true:", err)
}