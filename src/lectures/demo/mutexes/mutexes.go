package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1
	fmt.Printf("served iteration %d\n", iteration)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}

	hitCounter := Hits{}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		iteration := i
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Printf("Waiting for goroutines...\n\n")
	wg.Wait()

	fmt.Printf("Hit the server %d times\n", hitCounter.count)
}
