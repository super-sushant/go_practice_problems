package main

import "fmt"

func getNamesCount(userIds []string) (nameCountsMap map[string]map[string]int){
	nameCountsMap = make(map[string]map[string]int)
	return nameCountsMap
}

func main(){
	names:= []string{"jagdev","baldev","baba tillu","cheetah","sher","sher","baba tillu","sher"}

	namesCountsMap:= getNamesCount(names)

	// expect j:jagdev:1, b:baldev:1,baba tillu:2, s:sher:3, c:cheetah:1
	fmt.Printf("The returned user list is %v \n",namesCountsMap)
}