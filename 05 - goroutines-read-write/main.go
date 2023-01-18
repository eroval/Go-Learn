package main

import (
	"fmt"
	"sync"
)

func write(x *int, wg *sync.WaitGroup){
	*x++
	defer wg.Done()
}

func read(x *int, wg *sync.WaitGroup){
	fmt.Println(*x)
	defer wg.Done()
}

func main(){
	var data int = 0
	var wg sync.WaitGroup

	for i:=0; i<5; i++ {
		wg.Add(2)
		go write(&data, &wg)
		go read(&data, &wg)
	}
	wg.Wait()
	
	// Always will be correct in the end
	// but order of printing isn't guaranteed
	fmt.Println(data)
}