package encap_test

import (
	"testing"
	"fmt"
	"unsafe"
)

type Employee struct{
	Id string
	Name string
	Age int
  }


func TestCreateEmployeeObj(t *testing.T){
	e:=Employee{"0","bob",20}
	e1:=Employee{Name:"Mike",Age:30}
	e2:=new(Employee) //注意这里返回的引用/指针，相当于e:=&Employee{}
	e2.Id = "2"      //与其他主要编程语言的差异：通过实例的指针访问成员不需要使用->
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T",e)
	t.Logf("e is %T",&e)
	t.Logf("e2 is %T",e2)
}

func (e Employee) String() string{
	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
}

// func (e *Employee) String() string{
// 	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
// }

func TestStructOperations(t *testing.T){
	e:=Employee{"0","bob",20}
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	t.Log(e.String())
}