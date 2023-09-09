package main

import "fmt"

type student struct {
	name   string
	rollno int
	course string
	array1 []string
}

// Task # 02
type second_struct struct {
	id       int
	students []student
}

// Task # 01
// Passing struct as function argument
func structured(s student) {
	fmt.Printf("\n\t\t printing struct values")
	fmt.Println("")
	fmt.Println("Name: ", s.name)
	fmt.Println("Rollno: ", s.rollno)
	fmt.Println("Course: ", s.course)
	fmt.Println("Array: ", s.array1)
}

func main() {
	// Task # 01
	var s1 student
	var s2 student

	s1.name = "Huzaifa"
	s1.rollno = 20
	s1.course = "Blockchain"
	s1.array1 = []string{"1"}

	s2.name = "Abdullah"
	s2.rollno = 10
	s2.course = "SSD"
	s2.array1 = []string{"2"}

	fmt.Printf("\n\t\tStudent # 01\n")
	fmt.Println(s1)
	fmt.Printf("\n\t\tStudent # 02\n")
	fmt.Println(s2)
	// Task # 02
	// Calling structured function
	s3 := student{
		name:   "fahad",
		rollno: 15,
		course: "DLD",
		array1: []string{"3", "4", "5"},
	}
	structured(s3)
	// Task # 02
	sec1 := second_struct{
		id:       1,
		students: []student{s1, s2, s3},
	}
	fmt.Printf("\n\t\t Printing Second Structure\n")
	fmt.Println("")
	fmt.Printf("ID: %d", sec1.id)
	fmt.Println("")
	fmt.Println("struct: ", sec1.students)
}
