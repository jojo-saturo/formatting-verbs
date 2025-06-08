/*
package main

import "fmt"

func main() {
	// Intergers Formatting Verbs

	var Num int = 42

	fmt.Printf("%b \n", Num)  // Base 2
	fmt.Printf("%d \n", Num)  // Base 10
	fmt.Printf("%o \n", Num)  // Base 8
	fmt.Printf("%O \n", Num)  // Base 8 with leading zeros
	fmt.Printf("%x \n", Num)  // Base 16 Lowercase
	fmt.Printf("%X \n", Num)  // Base 16 Uppercase
	fmt.Printf("%#x \n", Num) // Base 16 Lowercase with leading 0x
	fmt.Printf("%4d \n", Num) // pad with spaces(width 4, right justified)
	fmt.Printf("%-4d \n", Num) // pad with spaces(width 4, left justified)
	fmt.Printf("%O4d \n", Num) // pad with zeros(width 4)
}
*/
// Float 64 Formatting Verbs
package main

import "fmt"

func main() {
	// Floats Formatting Verbs

	var Float float64 = 3.141592653589793

	fmt.Printf("%f \n", Float)   // Default format
	fmt.Printf("%.2f \n", Float) // Fixed point notation with 2 decimal places
	fmt.Printf("%e \n", Float)   // Scientific notation with 'e'
	fmt.Printf("%E \n", Float)   // Scientific notation with 'E'
	fmt.Printf("%g \n", Float)   // Default format or scientific notation if exponent is less than -4 or greater than 6
	fmt.Printf("%G \n", Float)   // Default format or scientific notation with 'E' if exponent is less than -4 or greater than 6
	fmt.Printf("%b \n", Float)   // Binary exponent
	fmt.Printf("%o \n", Float)   // Octal exponent
	fmt.Printf("%x \n", Float)   // Hexadecimal exponent
	fmt.Printf("%X \n", Float)   // Hexadecimal exponent with uppercase letters
	fmt.Printf("%#x \n", Float)  // Hexadecimal exponent with leading 0x and lowercase letters
	fmt.Printf("%4f \n", Float)  // pad with spaces(width 4, right justified)
	fmt.Printf("%-4f \n", Float) // pad with spaces(width 4, left justified)
	fmt.Printf("%O4f \n", Float) // pad with zeros(width 4)
}
