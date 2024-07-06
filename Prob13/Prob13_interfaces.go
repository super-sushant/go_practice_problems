package main

import "fmt"


func sendMessage (message msg) {
	println(message.sendMessage())
}

type msg interface {
	sendMessage() string
}

type birthdayMessage struct {
	birthdayDate string
	name string
}

type logMessage struct {
	errorMsg string
}

func (b birthdayMessage) sendMessage() string {
	return fmt.Sprintf("Hello %v happy birthday on %v",b.name,b.birthdayDate)
}

func (l logMessage) sendMessage() string {
	return fmt.Sprintf("Error: %v",l.errorMsg)
}



func main () {
	sendMessage(birthdayMessage{"Today","Chelsea Clinton"})
	sendMessage(logMessage{"Error on line 40"})
}