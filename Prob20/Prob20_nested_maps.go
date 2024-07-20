package main

import "fmt"

func getNamesCount(userIds []string) (nameCountsMap map[rune]map[string]int) {
	nameCountsMap = make(map[rune]map[string]int)

	for _, name := range userIds {
		if name == "" {
			continue
		}

		firstChar := rune(name[0])

		_, ok := nameCountsMap[firstChar]
		if !ok {
			// necessary value from retrieve is not referenced to the map
			nameCountsMap[firstChar] = make(map[string]int)
		}
		
		nameCountsMap[firstChar][name]++

	}

	return nameCountsMap
}

func main() {
	names := []string{"jagdev", "baldev", "baba tillu", "cheetah", "sher", "sher", "baba tillu", "sher"}

	namesCountsMap := getNamesCount(names)

	// expect j:jagdev:1, b:baldev:1,baba tillu:2, s:sher:3, c:cheetah:1
	fmt.Printf("The returned user list is %v \n", namesCountsMap)
}
