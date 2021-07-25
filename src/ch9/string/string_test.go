package string_test

import "testing"

func TestString(t *testing.T){
	var s string
	t.Log(s)
	s="hello"
	t.Log(len(s))       
	// s[1] = '3'    // string是不可变的byte slice
	s="\xE4\xB8\xA5"    //可以存储任何二进制数据
	t.Log(s)
	s="中"
	t.Log(len(s))

	c:= []rune(s)          //取出string转换为unicode
	t.Log(len(c))
	t.Logf("中 unicode %x",c[0])
	t.Logf("中 UTF8 %x",s)       // 输出16进制，用%x
}

func TestStringToRune(t *testing.T){
	s:="中华人民共和国"
	for _,c:= range s{
		t.Logf("%[1]c %[1]d",c)        //输出汉字、编码
	}
}