package main

import (
	"fmt"
	"sync"
)

type Number struct {
	x int
	y int
	xMu sync.Mutex	// used for locking X value when simultaneously writing to it
	yMu sync.Mutex  // used for locking Y value when simultaneously writing to it
}

func NewNumber(x int, y int) *Number {
	return &Number{x,y, sync.Mutex{}, sync.Mutex{}}
}

func (c *Number) AddX(x int){
	c.xMu.Lock()
	defer c.xMu.Unlock()
	c.x+=x
}

func (c *Number) AddY(y int){
	c.yMu.Lock()
	defer c.yMu.Unlock()
	c.y+=y
}

func (c* Number) Print(){
	fmt.Printf("x: %v\ny: %v\n", c.x, c.y)
}

func (c* Number) AddSimultaneously(x int, y int, number_of_times int){
	wg := sync.WaitGroup{}
	for i:=0; i<number_of_times; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			c.AddX(x)
		}()

		wg.Add(1)
		go func(){
			defer wg.Done()
			c.AddY(y)
		}()
	}

	wg.Wait()
	
}

func main(){
	fmt.Println("Cool")
	number:=NewNumber(5,5)
	number.AddX(5)
	number.Print()
	number.AddSimultaneously(5, 10, 100)
	number.Print()
}
