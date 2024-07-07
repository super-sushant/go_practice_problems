package main

import "fmt"

//variadic
func sum(nums ...int) int {
	var sum int
	for i:=0; i<len(nums); i++ {
		sum +=nums[i]

	}
	return sum
}


func main(){

	nums:= []int{1,2,3,4,5,6}

	// spread
	toral := sum(nums...)

	fmt.Println("The sum of the numbers is ",toral)
}