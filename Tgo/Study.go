package Tgo

type Human struct{
	Sex int
}
type Person struct{
	Human
	Name string
	Age int
}

func Add(a int,b int) int {
	return  a+b
}
