package string_fn

import(
	"testing"
	"strings"
	"strconv"
)

//strings包分割与拼接字符串
func TestStringFn(t *testing.T){
	s:="A,B,C"
	parts := strings.Split(s,",")
	for _,part:= range parts{
		t.Log(part)
	}
	t.Log(strings.Join(parts,"-"))
	
}

func TestIngConv(t *testing.T){
	s:=strconv.Itoa(10)
	t.Log("str" + s)
	if i,err:=strconv.Atoi("10");err==nil{
		t.Log(10+i)
	}
	
}