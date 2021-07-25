package interface_test

type Programmer interface{
	WriteHelloWorld() stirng
}

type GoProgrammer struct{

}

func (g *GoProgrammer) WriteHelloWorld() string{
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}