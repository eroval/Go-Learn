package Student

import (
	"errors"
	"fmt"
)

type Student struct{
	FirstName string
	LastName string
	FNumber string
}

func (student *Student) SetFirstName(name string) error {
	if(len(name)<=0){
		return errors.New("Invalid value")
	}
	student.FirstName = name
	return nil
}

func (student *Student) SetLastName(name string) error {
	if(len(name)<=0){
		return errors.New("Invalid value")
	}
	student.LastName = name
	return nil
}

func (student *Student) SetFNumber(number string) error {
	if(len(number)<=0){
		return errors.New("Invalid value")
	}
	student.FNumber = number
	return nil
}

func (student *Student) PrintStudent(){
	fmt.Println("Student:")
	fmt.Println("Name: \t\t", student.FirstName, student.LastName)
	fmt.Println("Faculty Number: ", student.FNumber)
}