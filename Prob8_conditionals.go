package main

import "fmt"

func main(){
	message:= "Hello papi muanyanyo"
	// messageLen:=  len(message)
	maxMessageLen:=20

	if messageLen:=len(message); messageLen<maxMessageLen {
		fmt.Println("Small message")
	}else if messageLen>maxMessageLen {
		fmt.Println("Large message")
	}else{
		fmt.Println("Perfect message")
	}
}