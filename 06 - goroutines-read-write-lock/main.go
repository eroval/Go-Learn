package main

import (
	"fmt"
	"sync"
)

func write(x *int, wg *sync.WaitGroup, access *sync.RWMutex){
	access.Lock()
	defer access.Unlock()

	*x++
	defer wg.Done()
}

func read(x *int, wg *sync.WaitGroup, access *sync.RWMutex){
	access.RLock()
	defer access.RUnlock()
	
	fmt.Println(*x)
	defer wg.Done()
}

func main(){
	var data int = 0
	var wg sync.WaitGroup
	var access sync.RWMutex


	for i:=0; i<5; i++ {
		wg.Add(2)
		go write(&data, &wg, &access)
		go read(&data, &wg, &access)
	}
	wg.Wait()
	
	fmt.Println(data)
}