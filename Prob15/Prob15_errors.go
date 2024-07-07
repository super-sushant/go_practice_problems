package main

import (
	"errors"
	"fmt"
)

func sendSms(msg string) (cost int, err error) {
	if len(msg) == 12 {
		return cost,errors.New("Wrong length")
	}
	cost = len(msg)
	return cost,err
}



func main() {
	cost, err:= sendSms("Hi my name 0")

	if(err!=nil){
		fmt.Printf("Couldn't send because %v \n",err)
		return
	}
	fmt.Println(cost)

}