package main

import "fmt"

func pointersmain() {
	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// fmt.Println(*firstName)

	firstName := "Arthur"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}
