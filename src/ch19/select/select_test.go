package select_test

import (
	"testing"
	"time"
	"fmt"
	// "errors"
)

func service() string{
	time.Sleep(time.Millisecond*100)
	return "Done"
}

func AsyncService() chan string{
	// retCh:=make(chan string)     //声明一个channel
	retCh:=make(chan string,1)     //声明一个buffer channel
	go func(){
		ret := service()
		fmt.Println("returned result.")
		retCh <-ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T){
	select {
	case ret:= <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond*100):
		t.Error("time out")
	}
}