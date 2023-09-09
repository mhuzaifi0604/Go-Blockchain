package main

import (
	"fmt"
)

// Creating struct for students
type students struct {
	name   string
	rollno int
	course string
	array1 []string
}

// function for creating a new student
func NewStudent(name string, rollno int, course string, array1 []string) *students {
	s1 := new(students)
	s1.name = name
	s1.rollno = rollno
	s1.course = course
	s1.array1 = array1
	return s1
}

// creating a structure containing list with datatype of prev structure
type studentlist struct {
	list []*students
}

// function to print the list of students added in student list
func (ls *studentlist) printstudents() {
	for i := 0; i < len(ls.list); i++ {
		fmt.Println("Student # ", i, " - ", *ls.list[i])
	}
}

// Function for creating a new student in student list
func (ls *studentlist) createstudent(name string, rollno int, course string, array1 []string) *students {
	// creating a new student
	st2 := NewStudent(name, rollno, course, array1)
	// appending the student in the student list
	ls.list = append(ls.list, st2)
	return st2
}

func main() {
	// Creating a new student list
	slist := new(studentlist)
	// Addeing stuents to teh list after creating them
	slist.createstudent("Huzaifa", 10, "BlockChain", []string{"1"})
	slist.createstudent("Moiz", 20, "CA", []string{"2"})
	// Pritning addresses of students in the list
	fmt.Println(*slist)
	fmt.Printf("\n\t\t Printing List of Students\n")
	// Printing data of each student in the list
	slist.printstudents()
}
