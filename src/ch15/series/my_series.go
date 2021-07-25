package series

import "fmt"

func init(){
	fmt.Println("init1")
}

// func square(n int) int{   //错，首字母必须大写。否则无法导入
// 	return n*n
// }
func Square(n int) int{   //对，首字母必须大写。否则无法导入
	return n*n
}

func GetFibonacciSerie(n int)[] int{
	ret := []int{1,1}
	for i:=2;i<n;i++{
		ret = append(ret,ret[i-2] + ret[i-1])
	}
	return ret
}