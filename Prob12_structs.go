package main

import "fmt"

type messageToSend struct {

}

func sendMessage(message messageToSend){
	fmt.Printf("Sending message: '%s' to: %v \n",message.message, message.phoneNumber)
}
func main () {
	message:="My name is Skylar White YO :3"
	phoneNumber:= 9876543211
	sendMessage (messageToSend{message,phoneNumber})
}