package main

import (
	"fmt"
	"runtime"
	"sync"
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

	sem := make(chan struct{}, poolSize)
	var wg sync.WaitGroup
	for _, d := range data {
		wg.Add(1)
		sem <- struct{}{}
		go func(d interface{}) {
			defer wg.Done()
			taskPerson(d)
			<-sem
		}(d)
	}
	wg.Wait()
}
