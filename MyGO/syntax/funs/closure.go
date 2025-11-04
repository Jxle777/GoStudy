package main

import "fmt"

func Closure(name string) func() string {
	return func() string {
		return "hello " + name
	}
}

// var age = 18
func Closure1( /*age int*/ ) func() string {
	name := "大明"
	age := 18
	return func() string {
		return fmt.Sprintf("Hello, %s, %d", name, age)
	}
}
