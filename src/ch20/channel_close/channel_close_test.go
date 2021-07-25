package channel_close

import(
	"testing"
	"fmt"
	"sync"
)

// 已知生产数据
// func dataProducer(ch chan int, wg *sync.WaitGroup){
// 	go func(){
// 		for i := 0;i <10;i++{
// 			ch <-i
// 		}
// 		wg.Done()
// 	}()
// }

// func dataReceiver(ch chan int, wg *sync.WaitGroup){
// 	go func(){
// 		for i := 0;i<10;i++{
// 			data := <-ch
// 			fmt.Println(data)
// 		}
// 		wg.Done()
// 	}()
// }

// func TestCloseChannel(t *testing.T){
// 	var wg sync.WaitGroup
// 	ch := make(chan int)
// 	fmt.Println(ch)
// 	wg.Add(1)
// 	dataProducer(ch,&wg)
// 	wg.Add(1)
// 	dataReceiver(ch,&wg)
// 	wg.Wait()
// }


func dataProducer(ch chan int, wg *sync.WaitGroup){
	go func(){
		for i := 0;i <10;i++{
			ch <-i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup){
	go func(){
		for i := 0;i<10;i++{
			if data,ok := <-ch;ok{
				fmt.Println(data)
			}else{
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T){
	var wg sync.WaitGroup
	ch := make(chan int)
	fmt.Println(ch)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Wait()
}