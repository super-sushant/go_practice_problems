package main

import "fmt"

func getCounts(userIds []string) (userCountMap map[string]int) {
	userCountMap = make(map[string]int)
	for _, userId := range userIds {
		count, ok := userCountMap[userId]
		if ok != true {
			userCountMap[userId] = 1
		} else {
			userCountMap[userId] = count + 1
		}
	}
	// get counts
	return userCountMap
}

func main() {
	userIds := []string{"user1", "user2", "user3", "user1", "user2", "user1"}

	countMap := getCounts(userIds)
	// expected map with user1:3 , user2:2 user3:1
	fmt.Printf("The count map of each user's count is %v. \n", countMap)

}
