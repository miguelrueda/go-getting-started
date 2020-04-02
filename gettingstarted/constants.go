package main

import "fmt"

const (
	first  = iota + 6
	second = 2 << iota
	pi     = 3.1415
)

const (
	third = iota
)

func constantsmain() {
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)

	fmt.Println(float32(c) + 1.2)

	fmt.Println(first, second)

	fmt.Println(third)
}
