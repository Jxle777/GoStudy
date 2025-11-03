package main

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	//he said:"Hello GO"
	println("he said:\"Hello GO\"")
	println("hello go")
	println(`
反引号优点
可以换行 反引号内不能再有反引号
这是新行 不能有转义符号
`)
	println("hello" + "go")               //go只能拼接字符串
	println("hello" + string(123))        //123会转为ascill码对应的符号
	println(fmt.Sprint("hello", 123))     //标准输出
	println(len("abc"))                   //3
	println(len("你好"))                    //6
	println(utf8.RuneCountInString("你好")) //2

}
