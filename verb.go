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
	fmt.Printf("%4d \n", Num)
	fmt.Printf("%-4d \n", Num)
	fmt.Printf("%O4d \n", Num)
}
