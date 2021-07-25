package fib

import (
	"fmt"
	"testing"
)

//1.go支持自动类型推导；2.go支持一个语句对多个变量赋值

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b     = 1
	// )
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 2
	b := 1
	a, b = b, a
	t.Log(a, b) //1,2
}
