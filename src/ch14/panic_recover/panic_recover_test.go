package panic_recover

import(
	"fmt"
	// "os"
	"testing"
	"errors"
)

func TestPanicVxExit(t *testing.T){

	// defer func(){
    //    fmt.Println("Finally!")
	// }()
	//测试recover
	defer func(){
		if err :=recover();err != nil{
			fmt.Println("recovered from",err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}