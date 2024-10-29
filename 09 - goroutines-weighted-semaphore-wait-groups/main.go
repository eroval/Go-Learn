package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

type Person struct {
	name string
	age  int
}

func taskPerson(d interface{}) {
	fmt.Println("Person ", d, " is doing the task asynchronously")
}

func main() {
	data := []interface{}{Person{name: "Lara", age: 15}, Person{name: "Kara", age: 15}}
	poolSize := runtime.NumCPU()

	fmt.Println("Number of people:", len(data))
	fmt.Println("CPU Threads:", poolSize)

	sem := semaphore.NewWeighted(int64(poolSize))
	var wg sync.WaitGroup
	for _, d := range data {
		wg.Add(1)
		sem.Acquire(context.Background(), 1)
		go func(d interface{}) {
			defer wg.Done()
			taskPerson(d)
			sem.Release(1)
		}(d)
	}
	wg.Wait()
}
