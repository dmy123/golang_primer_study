package fn_test

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)

func returnMultiValues() (int,int){
	return rand.Intn(10),rand.Intn(20)
}

func timeSpent(temp func(opt int) int) func(op int) int{
	// fmt.Println(opt)
	return func(n int) int{
		start := time.Now()
		ret := temp(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int)int{
	time.Sleep(time.Second*3)
	return op
}

func fastFun(op int)int{
	time.Sleep(time.Second*1)
	return op
}

func TestFn(t *testing.T){
	a,b:=returnMultiValues()
	t.Log(a,b)
	tsSF:=timeSpent(slowFun)
	t.Log(tsSF(7))
	t.Log(timeSpent(fastFun)(5))
}

func Sum(ops ...int) int{
	ret := 0
	for t,op := range ops{
		fmt.Println(t)
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

// func TestDefer(t *testing.T){
// 	defer func(){
// 		t.Log("Clear resources")
// 	}()
// 	t.Log("Started")

// 	panic("Fatal error")    //defer仍会执行
// 	t.Log("after panic")
// }

func Clear(){
	fmt.Println("clear.")
}

func TestDefer(t *testing.T){
	defer Clear()
	t.Log("start")
	panic("fatal error")
}