package main

import (
	"fmt"
	"time"
)

func main(){
	test()
}

func saveBackups(snapshotTicker,saveAfter <-chan time.Time) {
	for {

	select {
		case <-saveAfter: 
			saveSnapshot()
			return
		case <-snapshotTicker: 
			takeSnapshot()
		
		default :
			waitForData()
			time.Sleep(500*time.Millisecond)
	}
}
}

func waitForData(){
	fmt.Println("Nothing to do, waiting...")

}

func takeSnapshot(){
	fmt.Println("Taking a backup snapshot ...")
}

func saveSnapshot(){
	fmt.Println("All backups saved!")
}

func test(){
	snapshotTicker := time.Tick(800* time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("==================>")
}