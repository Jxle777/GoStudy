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

// 子切片
func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3] //左闭右开区间 容量从 4 开始数后面的都算
	fmt.Printf("s2: %v len=%d, cap=%d ", s2, len(s2), cap(s2))

	//容量就是 start 开始往后([start:end])，包括原本 s1 的底层数组元素的个数
	s3 := s1[2:] //从6开始算 容量未定义最大为3 长度为3
	fmt.Printf("s3: %v len=%d, cap=%d ", s3, len(s3), cap(s3))

	s4 := s1[:3] //从2开始算 长度为3 容量为5
	fmt.Printf("s4: %v len=%d, cap=%d ", s4, len(s4), cap(s4))
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[2:3]
	fmt.Printf("s2: %v len=%d, cap=%d ", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s2: %v len=%d, cap=%d ", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v len=%d, cap=%d ", s1, len(s1), cap(s1))

	s2 = append(s2, 199)
	fmt.Printf("s2: %v len=%d, cap=%d ", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v len=%d, cap=%d ", s1, len(s1), cap(s1))

	s2[1] = 199
	fmt.Printf("s2: %v len=%d, cap=%d ", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v len=%d, cap=%d ", s1, len(s1), cap(s1))
}
