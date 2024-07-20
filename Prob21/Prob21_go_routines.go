package main

import "fmt"


func waitForDbs(numDBs int,dbChan chan struct{}) {
	
}


func getDatabasesChannel(numDbs int) chan struct{} {
	ch := make(chan struct{})

	go func() {
		for i:=0; i< numDBs; i++ {
			ch <-struct{}{}
			fmt.Printf("Data %v is online \n",i+1)
		}
	}()

	return ch
}



func main (){
	dbChan:=getDatabasesChannel(10)

	waitForDbs(10,dbChan)
}