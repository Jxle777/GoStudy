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

func UserSumInt64() {
	s1 := []int{1, 2, 3}
	res := SumInt64(s1)
	println(res)
}

func SumInt(vals[int]) int {
	var res int
	for _, val := range vals {
		res += val
	}
	return res
}

func SumInt64(vals []int64) int64 {
	var res int64
	for _, val := range vals {
		res += val
	}
	return res
}

func UserKey() {
	m := map[string]int{
		"key1": 123,
	}
	keys := Keys(m)
	println(keys)

} //map 键和值的类型一定要对的上

func Keys(m map[string]any) []string {
	return []string{"hello"}
}
