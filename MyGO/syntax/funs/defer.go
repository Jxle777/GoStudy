package main

// （栈）后进先出
func Defer() {
	defer func() {
		println("第一个defer")
	}()

	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

// 实际运用举例
//func Query() {
//	db, _ := sql.Open("", "")
//	defer db.Close()
//	db.Query("")
//}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
} //defer不能改:=的值 输出0

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
} //输出1
