package main

import (
	"fmt"
	"sync"
	"time"
)

// var mex = sync.Mutex{}
var mex = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
var results = []string{}

func main() {
	fmt.Println("Go Routine")
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\ntotal execution time: %v\n", time.Since(t0))
	fmt.Println("results are :", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep((time.Duration(delay) * time.Millisecond))
	fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	mex.Lock()
	results = append(results, result)
	mex.Unlock()
}

func log() {
	mex.RLock()
	fmt.Println("the current results are: ", results)
	mex.RUnlock()
}
