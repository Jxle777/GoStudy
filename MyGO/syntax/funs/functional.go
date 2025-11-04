package main

func Functional4() {
	println("Hello Functional4")
}

func Functional5(age int) {
	println(age)
}

func UseFunctional4() {
	myFunc := Functional4
	myFunc()
	myFunc5 := Functional5
	myFunc5(18)
}

func Functional6() {
	//新定义了一个方法赋值给了 fn
	fn := func() string {
		return "Hello"
	}

	fn()
}
