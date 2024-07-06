package main 

import "fmt"


func add(x string,y string) string {
	x = "sup" 
	return x+y
}

func sub(x,y string) string {
	return x+y
}

func main() {
	x:= "Hello"
	fmt.Println(add(x,"Bitch"))
	fmt.Println(sub(add(x,"Ugly"),"Bitch"))
	fmt.Println(x)
}

var f func (int int) int 