package main

func Switch(status int) {
	switch status { //带了status为对应类型的值
	case 0:
		println("init")
	case 1:
		println("running")
	default: //可有可无
		println("unknow")
	}
	println("都没有执行直接执行这句")
}

func SwitchBool(age int) {
	switch { //不带为bool类型的值
	case age >= 18:
		println("成年人")
	case age >= 12:
		println("青年人")
	default:
		println("")
	}
}
