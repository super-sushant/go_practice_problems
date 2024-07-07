package main

import "fmt"



func main(){
	name:= []string{"Shakti","Bhakti","Rakti","Nakti"}
	numbers:= []int{12,13,14}
	userDict := getUserMap(name,numbers)

	fmt.Print(userDict)
}

func getUserMap(names []string, numbers []int) (myMap map[string]int) {
	myMap= make(map[string] int)

	for i,name :=range names {
		if len(numbers)>i {
			myMap[name]= numbers[i]
		}else {
			myMap[name]=0
		}
	}
	
	return myMap
}