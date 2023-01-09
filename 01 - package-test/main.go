package main

import (
	"fmt"
	"package-test.com/app/local_package"
)

func main() {
	var x,y int = 5, 10
	
	fmt.Println(local_package.AddValues(x, y))
}
