package main

func Bool() {
	var a bool = true
	var b bool = false
	var c bool = a || b //true
	println(c)
	var d bool = a && b //flase
	println(d)
	var e bool = !a
	println(e)
	var f bool = !(a && b) //相当于!a||!b
	println(f)
	var g bool = !(a || b) //!a&&!b
	println(g)

}
