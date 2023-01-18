package main

import (
	"fmt"
	"sync"
)

func myfunction(x *int, wg *sync.WaitGroup){
	*x++
	defer wg.Done()
}

func main(){
	var data int = 0
	var wg sync.WaitGroup

	for i:=0; i<5; i++ {
		wg.Add(1)
		go myfunction(&data, &wg)
	}
	wg.Wait()
	
	fmt.Println(data)
}