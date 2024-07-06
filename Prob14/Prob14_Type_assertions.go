package main

import "fmt"

type email struct {
	toAddress string
	cost      float64
}

type sms struct {
	toPhoneNumber uint
}

type expense interface {
	getCost() float64
}

func (e email) getCost() float64 {
	return e.cost
}

func (message sms) getCost() float64 {
	return float64(message.toPhoneNumber)
}

// type assert to return to also print emailaddress and phoneNumber
func getExpense(e expense) {
	fmt.Println(e.getCost())
}

func main() {
	myEmail := email{
		toAddress: "sabrinacarpenter@majdoor.com",
		cost: 6000,
	}

	mySms:= sms{
		toPhoneNumber: 9211420121,
	}

	getExpense(myEmail)
	getExpense(mySms)

}
