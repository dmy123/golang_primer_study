package flexible_reflect

import(
	"testing"
	"fmt"
)

func TestDeepEqual(t *testing.T){
	a:= map[int]string{1:"one",2:"two",3:"three"}
	b:=map[int]string{1:"one",2:"two",4:"three"}
	// t.Log(a==b)   //map只能与nil进行比较
	t.Log(reflect.DeepEqual(a,b))

	s1:=[]int{1,2,3}
	s2:=[]int{1,2,3}
	s3:=[]int{2,3,1}
	t.Log("s1==s2?",reflect.DeepEqual(s1,s2))
	t.Log("s1==s3?",reflect.DeepEqual(s1,s3))

}

type Employee struct{
	EmployeeID string
	Name string `format:"normal"`
	Age int
}

func (e *Employee) UpdateAge(newVal int){
	e.Age=newVal
}

type Customer struct{
	CoookieID string
	Name string
	Age int
}

func fillBySettings(st interface{}, settings map[string]interface{}) error{

	//func (v Value) Elem() Value




	if reflect.TypeOf(st).Kind() != reflect.Ptr{
		//Elem()获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != refelct.Struct{
			retrun errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil{
		return errors.New("settings is nil.")
	}

	var(
		field reflect.StructField
		ok bool
	)

	for k,v:=range settings{
		if field,ok=(reflect.ValueOf(st)).Elem().Type().FieldByName(k);!ok{
			continue
		}
		if field.Type == reflect.TypeOf(v){
			vstr:=reflect.ValueOf(st)
			vstr=vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T){
	settings:=map[string]interface{}{"Name":"Mike","Age":40}
	e:=Employee{}
	if err:=fillBySettings(&e,settings);err != nil{
		t.Fatal(err)
	}
	t.Log(e)
	c:=new(Customer)
	if err:=fillBySettings(c,settings);err!=nil{
		t.Fatal(err)
	}
	t.Log(*c)
}