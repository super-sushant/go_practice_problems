package main

import (
	"fmt"
	"time"
)

func main(){
	test()
}

func saveBackups(snapshotTicker,saveAfter <-chan time.Time) {
	
}

func waitForData(){
	fmt.Println("Waiting for data....")
}

func takeSnapshot(){
	fmt.Println("Taking a backup snapshot ...")
}

func saveSnapshot(){
	fmt.Println("Nothing to do, waiting...")
}

func test(){
	snapshotTicker := time.Tick(800* time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("==================>")
}