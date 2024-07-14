package main

import "fmt"

func getCounts(userIds []string) map[string]int {
	// get counts
	return make(map[string]int)
}

func main() {
	userIds := []string{"user1", "user2", "user3", "user1", "user2", "user1"}

	countMap := getCounts(userIds)
	// expected map with user1:3 , user2:2 user3:1
	fmt.Printf("The count map of each user's count is %v. \n",countMap)

}
