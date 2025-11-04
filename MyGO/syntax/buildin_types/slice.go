package main

import "fmt"

func Slice() {
	s1 := []int{9, 8, 7}
	fmt.Printf("a1: %v len=%d, cap=%d ", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) //第一个是长度 第二个是容量
	fmt.Printf("s2: %v len=%d, cap=%d ", s2, len(s2), cap(s2))

	s3 := make([]int, 4) //长度和容量都是4
	fmt.Printf("s3: %v len=%d, cap=%d ", s3, len(s3), cap(s3))

	s4 := make([]int, 0, 4)
	s4 = append(s4, 1)
	fmt.Printf("s4: %v,len=%d, cap=%d ", s4, len(s4), cap(s4))
}
