package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

// var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nTotal results are: %v", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	// var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from the databse is: ", dbData[i])
	// m.Lock()
	// results = append(results, dbData[i])
	// m.Unlock()
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}