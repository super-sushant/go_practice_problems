package main

import "fmt"

type messageToSend struct {
	message string
	phoneNumber int
}

func sendMessage(message messageToSend){
	fmt.Printf("Sending message: '%s' to: %v \n",message.message, message.phoneNumber)
}
func main () {
	phoneNumber:="My name is Skylar White YO :3"
	message:= 9876543211
	sendMessage (messageToSend{
		phoneNumber,
		message,
	})
}