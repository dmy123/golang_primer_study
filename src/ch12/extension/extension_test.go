package extension

import (
	"testing"
	"fmt"
)
// //1.结果...  Chao
// type Pet struct{

// }

// func (p *Pet) Speak(){
// 	fmt.Print("...")
// }

// func (p *Pet)SpeakTo(host string){
// 	p.Speak()
// 	fmt.Println(" ",host)
// }

// type Dog struct{
// 	p *Pet
// }

// func (d *Dog) Speak(){
// 	fmt.Println("Wang!")
// }

// func (d *Dog)SpeakTo(host string){
// 	d.p.SpeakTo(host)
// }

// func TestDog(t *testing.T){
// 	dog:= new(Dog)
// 	dog.SpeakTo("Chao")
// }

// //2.结果Wang!
// //  Chao
// type Pet struct{

// }

// func (p *Pet) Speak(){
// 	fmt.Print("...")
// }

// func (p *Pet)SpeakTo(host string){
// 	p.Speak()
// 	fmt.Println(" ",host)
// }

// type Dog struct{
// 	p *Pet
// }

// func (d *Dog) Speak(){
// 	fmt.Println("Wang!")
// }

// func (d *Dog)SpeakTo(host string){
// 	d.Speak()
// 	fmt.Println(" ",host)
// }

// func TestDog(t *testing.T){
// 	dog:= new(Dog)
// 	dog.SpeakTo("Chao")
// }

//3.结果...  Chao
//说明这种内嵌结构类型完全不能当成继承，无法访问子类方法数据，即不支持重载
type Pet struct{

}

func (p *Pet) Speak(){
	fmt.Print("...")
}

func (p *Pet)SpeakTo(host string){
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct{
	Pet
}

func (d *Dog) Speak(){
	fmt.Println("Wang!")
}


func TestDog(t *testing.T){
	dog:= new(Dog)
	dog.SpeakTo("Chao")
}