package main

import "fmt"

func main () {
	firstname, _ := getNames()

	fmt.Printf("Hi %v ass boi",firstname)

}

func getNames()(string,string){
	return "Bitch", "lasagna"
}