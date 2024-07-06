package main

import "fmt"

func main() {
	x := 10

	x=inc(x)

	fmt.Println(x,"Ouside increment")


}

func inc(x int)int {
	x = x * 2
	fmt.Println(x,"Inside increment")
	return x
}
