package main

func Map() {
	m1 := map[string]int{
		"key1": 1,
	}
	m1["hello"] = 345

	m2 := make(map[string]int, 12)
	m2["key2"] = 12

	val, ok := m1["大明"]
	if ok {
		//有这个键值对
		println(val)
	}

	val = m1["小明"] //在 Go 中，当访问一个 map 中不存在的键时，返回该值类型的 零值。
	println("小明对应的值：", val)

	delete(m1, "key1")
}
