package main

var Gloal = "全局变量"
var internal = "包内变量，私有变量"
var (
	first  string = "1"
	second string = "2"
	aa            = "hello"
)

func main() {
	var a int = 123 //go会类型推断
	println(a)

	var b = 234
	println(b)

	var c uint = 456
	println(c)

	var (
		d string = "aaa"
		e int    = 123
	)
	println(d, e)

	f := 123
	println(f)

	var aa int = 123 //赋值覆盖，最好避免这样写
	println(aa)
}
