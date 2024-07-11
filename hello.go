package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, my name is James.")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
}

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
