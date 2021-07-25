package first_response

import(
	"testing"
	"fmt"
	"time"
	"runtime"
)
//test1
// func runTask(id int) string{
// 	time.Sleep(10*time.Millisecond)
// 	return fmt.Sprintf("The result is from %d",id)
// }

// func FirstResponse() string{
// 	numOfRunner := 10
// 	ch := make(chan string)
// 	for i:=0;i < numOfRunner;i++{
// 		go func(i int){
// 			ret := runTask(i)
// 			ch <-ret
// 		}(i)
// 	}
// 	return <-ch
// }

// func TestFirstResponse(t *testing.T){
// 	t.Log("Before",runtime.NumGoroutine())  //2
// 	t.Log(FirstResponse())//1
// 	time.Sleep(time.Second*1)
// 	t.Log("After:",runtime.NumGoroutine())   //11
// }

//test2,避免资源浪费
func runTask(id int) string{
	time.Sleep(10*time.Millisecond)
	return fmt.Sprintf("The result is from %d",id)
}

func FirstResponse() string{
	numOfRunner := 10
	ch := make(chan string,numOfRunner)   //用buffer channel避免资源泄露
	for i:=0;i < numOfRunner;i++{
		go func(i int){
			ret := runTask(i)
			ch <-ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T){
	t.Log("Before",runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second*1)
	t.Log("After:",runtime.NumGoroutine())
}