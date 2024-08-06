//go:build ignore

// Every Go program must start with a package declaration
package main

import (
	"fmt"
)

// Functions
func returnMinimalistic(name string) string {
	return "Hello " + name
}
func returnSpecify(name string) string {
	var greet string = "Hello " + name
	return greet
}

func onlyReturnTypeSpecify(name string) string {
	var greet string = "Hello " + name
	return greet
}

func multiReturn(a int, b int) (int, int, bool) {
	return a + b, a - b, true
}

type Patient struct {
	ptr  *Patient
	name string
	age  uint8
	rank rune
}

type Car struct {
	color   string
	seats   int
	price   int
	company struct {
		name  string
		model string
	}
}

func (car Car) rent() int {
	return car.seats * 20
}

func main() {
	/*Spaces are always added between operands and a
	newline is appended. It returns the number of
	bytes written and any write error encountered.*/
	fmt.Println("Hello, World!")

	const profit = 20
	fmt.Printf("my profit is %v%%\n", profit)

	// define variables
	const Root3 float32 = 1.71
	const Euler float32 = 2.71828
	println(Root3, Euler)

	const (
		X float32 = 2.45
		Y float32 = 1.67
		Z float32 = 3.78
	)
	print(X, Y, Z)

	var integer uint8 = 255
	var decimal float32 = 3.14
	var character rune = 'a'
	var str string = "Jhon"
	var boolean bool = true
	var pointer *uint8 = &integer

	/*
		%v - default format for the value
		%T - type of the value
		%t - boolean (true or false)
		%d - decimal integer
		%+d- Base 10 and always show sign
		%4d- Pad with spaces (width 4, right justified)
		%-4d-Pad with spaces (width 4, left justified)
		%04d-Pad with zeroes (width 4
		%c - charachter representation
		%b - binary representation
		%o - octal representation
		%O - Base 8, with leading 0o
		%x - hexadecimal representation (with lowercase letters)
		%X - hexadecimal representation (with uppercase letters)
		%#x- Base 16, with leading 0x
		%s - Prints the value as plain string
		%q - double-quoted string safely escaped with Go syntax
		%f - floating-point representation
		%6.2f	Width 6, precision 2
		%8s- Prints the value as plain string (width 8, right justified)
		%-8s Prints the value as plain string (width 8, left justified)
		%e - scientific notation (e.g., 1.23e+04)
		%E - scientific notation (e.g., 1.23E+04)
		%g - %e for large exponents, %f otherwise
		%G - %E for large exponents, %f otherwise
		%p - pointer representation
	*/

	// specific print functionality
	fmt.Printf("%d\t\t%T\n%f\t\t%T\n%c\t\t%T\n%q\t\t%T\n%t\t\t%T\n%p\t\t%T\n", integer, integer, decimal, decimal, character, character, str, str, boolean, boolean, pointer, pointer)

	// if else statement
	var ifVar, elifvar, innerIf = true, false, true
	if !ifVar {
		print("if condition satisfy")
		if innerIf {
			print("inner if condition satisfy")
		} else {
			print("inner else condition satisfy")
		}
	} else if elifvar {
		print("else-if condition satisfy")
	} else { //this should be on same line
		print("else condition satisfy")
	}

	//Array
	var arr2D = [2][2]bool{}
	var array = [5]int{1, 2, 3, 4}
	var semiInitialiseArr = [...]rune{0: 'a', 2: 'z'}
	var uninitiateArr = [5]string{}
	arr2D[0][0] = true

	fmt.Println(arr2D, array, semiInitialiseArr, uninitiateArr)
	fmt.Println(len(arr2D))

	// Slices LSF
	// Using the []datatype{values} format
	var slice0 = []int{} // 0 length 0 capacity
	var slice1 = []string{"length", "is", "3"}
	fmt.Println(len(slice0), cap(slice0))
	fmt.Println(len(slice1), cap(slice1))

	fmt.Println("slice0 before append", slice0)
	slice0 = append(slice0, 1)
	fmt.Println("slice0 after append", slice0)

	//Create a Slice From an Array
	var sliced_array = array[1:3]
	fmt.Println(sliced_array)

	//Create a Slice With The make() Function
	// Note: If the capacity parameter is not defined, it will be equal to length.
	scratch_slice := make([]int, 5, 10)

	fmt.Println(len(scratch_slice), cap(scratch_slice))

	//Switch Statement
	//All the case values should have the same type as the switch expression
	price := 3
	switch price {
	case 1, 2, 3:
		fmt.Println("price is less than 3")
	default:
		fmt.Println("price is greater to 3")
	}

	// loops
	for i := 0; i <= 5; i++ {
		if i < 4 {
			fmt.Println(i)
			continue
		} else {
			break
		}
	}

	//Range Keyword LSF
	// returns both the index and the value.
	fmt.Println("2D array")
	for index, value := range arr2D {
		fmt.Printf("%v\t%t\n", index, value)
	}
	// while loop using for loop
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 2 {
			break
		}
	}

	// function call
	fmt.Println(returnMinimalistic("vaibhav"))
	var (
		add, sub int
		state    bool
	)
	add, sub, state = multiReturn(5, 4)
	fmt.Printf("a:%d\t\tb:%d\n", 5, 4)
	fmt.Printf("Addition:%d\tSubstraction:%d\n", add, sub, state)

	//Struct
	var rick Patient = Patient{name: "Rick Sanchez", age: 100, rank: 'A'}
	var morty Patient = Patient{ptr: &rick, name: "Morty sanchez", age: 18, rank: 'D'}
	fmt.Println(morty)

	myCar := Car{
		color: "Blue", seats: 5, price: 30000,
		company: struct {
			name  string
			model string
		}{name: "Toyota", model: "Camry"},
	}
	fmt.Println("car rent is:", myCar.rent())

	// Maps
	dict := make(map[string]int)
	dict["apple"] = 10
	dict["orange"] = 5
	fmt.Println(dict)
	delete(dict, "orange")
	fmt.Println(dict)

	var variable, ok = dict["apple"]
	if ok {
		fmt.Printf("%T\n", variable)
	}

	var elements map[string]string = map[string]string{"H": "Hydrogen", "He": "Helium", "Li": "Lithium"}
	fmt.Println(elements)

	// Closure
	var math func(int, int, string) int = func(opt1 int, opt2 int, operation string) int {
		var addition func(int, int) int = func(a int, b int) int {
			return a + b
		}
		var substraction func(int, int) int = func(a int, b int) int {
			return a - b
		}
		var product func(int, int) int = func(a int, b int) int {
			return a * b
		}

		switch operation {
		case "add":
			return addition(opt1, opt2)
		case "sub":
			return substraction(opt1, opt2)
		case "prod":
			return product(opt1, opt2)
		default:
			return 0
		}
	}
	var operation string = "add"
	fmt.Println(math(3, 4, operation))

	// Recursion

	// Defer, Panic & Recover

	// define array, dictionary, hashmap, linkedlist
	// array, dictionary, hashmap methods
	// read/write system files
	// create and import packages

}
