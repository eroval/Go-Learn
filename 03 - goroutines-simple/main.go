package main

import (
	"fmt"
	"time"
)

func myfunction(x *int){
	*x++
}

func main(){
	var data int = 0
	for i:=0; i<5; i++ {
		go myfunction(&data)
	}
	time.Sleep(time.Second)
	fmt.Println(data)
}