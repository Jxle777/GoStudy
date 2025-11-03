package main

func main() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a - b)
	println(a * b)
	if b != 0 {
		println(a / b)
		println(a % b)
	}
	var c float64 = 3.14
	println(a + int(c)) //go只能同类型加减乘除
	var d int32 = 456
	var e int64 = 123
	//println(d + e)//同样报错
	println(d + int32(e))
	//math.Ceil()//取最近整数
	String()
}
