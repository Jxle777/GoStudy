package main

func YourName(name string, aliases ...string) {
	//	aliases 是一个切片
}

// 不定参数
func CallYourName() {
	YourName("大明")
	YourName("大明", "小明")
	YourName("大明", "小明", "daming")
	aliases := []string{"大明", "小明"}
	YourName("大明", aliases...)

}
