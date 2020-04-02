package main

import "fmt"

func structsmain() {
	type user struct {
		ID        int
		Firstname string
		Lastname  string
	}

	var u user
	u.ID = 1
	u.Firstname = "arthur"
	u.Lastname = "dent"
	fmt.Println(u)

	u2 := user{
		ID:        2,
		Firstname: "miguel",
		Lastname:  "rueda",
	}
	fmt.Println(u2)
}
