package main

import "fmt"

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i:=0; i<numDBs; i++ {
		<-dbChan
	}
}

func getDatabasesChannel(numDbs int) chan struct{} {
	ch := make(chan struct{})

	go func(numDBs int) {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Data %v is online \n", i+1)
		}
	}(numDbs)

	return ch
}

func main() {
	dbChan := getDatabasesChannel(10)

	waitForDbs(10, dbChan)
}
