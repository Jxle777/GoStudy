package main

func main() {
	//name, age := Func10() //发起方法调用
	//println(name, age)
	//
	//name1, _ := Func10() //_是忽略那个值
	//println(name1)
	////使用 := 的前提 左边必须至少有一个新变量
	//name1, name2 := Func10()
	//println(name1)
	//println(name2)
	//
	//Func6("hello", "go") //传参数
	//Recursive(10)//递归调用
	//UseFunctional4()
	//Functional8()
	//fn := Closure("Daming")
	//fn 其实已经从 Closure 里面返回了
	//但是我 fn 还要用到“大明”
	//println(fn())
	//Defer()
	//DeferClosure()
	//DeferClosureV1()
	//println(DeferReturn())
	//println(DeferReturnV1())
	DeferClosureLoopV1()
	println("\n")
	DeferClosureLoopV2()
	println("\n")
	DeferClosureLoopV3()
}
