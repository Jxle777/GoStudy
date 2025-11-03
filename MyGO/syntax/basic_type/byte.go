package main

import "fmt"

func Byte() {
	var a byte = 'a'
	println(a)                    //打印对应的ascll码值
	fmt.Printf("%c\n", a)         // ② 按字符输出
	println(fmt.Sprintf("%c", a)) // ③ 格式化为字符串再打印
	var str string = "this is string"
	var bs []byte = []byte(str)
	bs[0] = 'T' //切片不更改值
	println(str)
}
