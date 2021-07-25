package groutine_test

import (
	"fmt"
	"testing"
	// "time"
)

func TestGroutine(t *testing.T){
	for i:=0;i<10;i++{
		go func (i int){
			fmt.Println(i)
		}(i)
	}

}

// func TestGroutineShared(t *testing.T){
// 	for i:=0;i<10;i++{
// 		go func (){
// 			fmt.Println(i)
// 		}()
// 	}
// 	// time.Sleep(time.Millisecond*100)
// }