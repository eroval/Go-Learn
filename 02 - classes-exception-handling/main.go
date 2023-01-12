package main

import (
	"os"
	"classes-exception-handling.com/app/Student"
)

// Mimic usual behaviour
func run() int{
	student := Student.Student{}
	exception := student.SetFirstName("cool")
	if (exception!=nil){
		return 1
	}
	exception = student.SetLastName("Go might not be that cool")
	if (exception!=nil){
		return 2
	}
	exception = student.SetFNumber("095108")
	if (exception!=nil){
		return 3
	}
	student.PrintStudent()
	return 0
}

func main() {
	os.Exit(run())
}
