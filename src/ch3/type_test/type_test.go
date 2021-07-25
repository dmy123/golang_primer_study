package type_test

import "testing"

type MyInt int64

//总：1.go不支持隐式类型转换，包括别名；2.go可以用指针，但不支持指针运算；3.go中字符串是值类型，初始化空字符串是长度为0空字符串，而不是nil

func TestImplicit(t *testing.T){
	var a int = 1
	var b int64
	// b=a   //go不支持隐式类型转换
	// b=int64(a)
	var c MyInt
	c=b            //go不支持隐式类型转换，包括别名
	t.Log(a,b,c)
}

func TestPoint(t *testing.T){
	a:=1
	aPtr:= &a
	t.Log(a,aPtr,*aPtr)
	t.Logf("%T %T",a,aPtr)   //获取类型
}

func TestString(t *testing.T){
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
	// if s == ""{            //判断空字符串条件

	// }
}