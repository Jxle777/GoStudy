package main

import "fmt"

func Array() {
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: len=%d, cap=%d ", a1, len(a1), cap(a1))

	println("\n")

	a2 := [3]int{9, 8}
	fmt.Printf("a2: len=%d, cap=%d ", a2, len(a2), cap(a2))

	println("\n")

	var a3 [3]int
	fmt.Printf("a3: len=%d, cap=%d ", a3, len(a3), cap(a3))

	//a3 = append(a3,12)
	println(a3[2])
}
