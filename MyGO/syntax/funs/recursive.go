package main

// 递归使用不当，就有可能会出现stack overflow(go默认栈为2KB)
func Recursive(n int) {
	if n > 10 {
		return
	}
	Recursive(n + 1)
}
