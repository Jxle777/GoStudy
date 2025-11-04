package main

func IfOnly(age int) {
	if age >= 18 {
		println("成年了")
	}
}

func IfelseIf(age int) string {
	if age >= 18 {
		println("成年了")
	} else if age >= 12 {
		println("青年")
	} else {
		println("小孩")
	}
	return ""
}

func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "太远了"
	} else if distance > 60 {
		return "有点远"
	} else {
		return "还挺好"
	}

	//if xx, err := xxx; err != nil {
	//
	//}之后会用到的写法

	//println(distance) 报错，作用域很小
}
