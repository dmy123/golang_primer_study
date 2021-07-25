package conditon_test

import "testing"

func TestSwitchMultiCase(t *testing.T){
	for i:=0;i < 5;i++{
		switch i{
		case 0,2:
			t.Log("even")
		case 1,3:
			t.Log("odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitch(t *testing.T){
	for i:=0;i < 5;i++{
		switch {
		case i%2==0:
			t.Log("even")
		case i%2==1:
			t.Log("odd")
		default:
			t.Log("unknow")
		}
	}
}

func TestIf(t *testing.T){
	if a:=2;a<0{
		t.Log("yes")
	}else{
		t.Log("no")
		t.Log(a)
	}
	// t.Log(a)
}