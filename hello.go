package main

import "fmt"

// commented line
/*
	commented section
*/

/*
Variables
In Go are statically typed, once they type is set it cannot be changed.


Integers
int - integers have no decimal component
unint8 - unsigned integer, cannot be negative only 0 or greater.
unint8 - unsigned integer with a length of 8 bits
int64 - signed integer of 64 bits, we can use 8, 16, 32 and 64 though we mostly just use int.

Floating Point
float32 - floats can be two sizes, 32 bit or 64bit. Single precision.
float64 - floats are inexact, close but not exact, float64 is used more often. Double precision.

Operators
+ - addtion
- - subtraction
* - multiplication
/ - division
% - remainder


Strings
Use double quotes "" or single quotes '' to create strings.
"" - double quotes cannot contain new lines, but allow special escape sequences.
	\n - becomes newline
	\t - becomes tab
'' - single quotes can have new lines, but not special escape sequences.

len("Hello World") - finds string length.
"Hello World"[1] - returns the byte value "0101" of the individual character at position 1 "e".
"Hello"+"World" - concatenates two strings together.


Booleans
&& - and
|| - or
! - not
*/

// new type
type person struct {
	// lowercase not visible outside the package
	// uppercase visible outside the package
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says Good morning.")
}

type angryPerson struct {
	person
	hasCoffe bool
}

func (aP angryPerson) speak() {
	if aP.hasCoffe == true {
		fmt.Println("Good morning.")
	} else {
		fmt.Println("Do not speak to me I have not had my coffee yet.")
	}
}

type human interface {
	speak()
}

func haveAChat(h human) {
	h.speak()
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, my name is James.")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	// Variable declaration
	var x int
	x = 2
	fmt.Println(x)

	// Variable short declaration
	y := 1
	fmt.Printf("%T", y)

	// List
	xSlice := []int{1, 1, 2, 4, 6, 12}
	fmt.Println(xSlice)

	// Map, key:value
	yMap := map[string]int{
		"James": 19,
		"Laura": 90,
	}
	fmt.Println(yMap)

	// if we use all values of the type we do not need to declare subtypes.
	p1 := person{"James", "Tynan"}
	fmt.Println(p1)
	p1.speak()

	p2 := person{}
	p2.lname = "Smith"
	p2.fname = "Laura"
	fmt.Println(p2)

	p3 := angryPerson{
		person:   person{"Grump", "Butt"},
		hasCoffe: false,
	}
	haveAChat(p3)
	p3.hasCoffe = true
	haveAChat(p3)
}
