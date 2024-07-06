package main

import "fmt"

func main (){
	const main = "Ryan Gosling"
	const openRate = 30.5

	// msg := "Hi Ryan Gosling, your open rate is 30.50 percent"
	msg := fmt.Sprintf("Hi %d, your open rate is %d percent",main,openRate)
	fmt.Println(msg)
}