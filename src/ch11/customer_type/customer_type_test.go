package customer_type

import "testing"

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int)int{
	time.Sleep(time.Second*3)
	return op
}

func TestFn(t *testing.T){
	tsSF:=timeSpent(slowFun)
	t.Log(tsSF(7))
}