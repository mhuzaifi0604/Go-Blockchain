package main
import "fmt"

type students struct{
	name string
	rollno int
	course string
	array1 []string
}

func NewStudent(name string, rollno int, course string, array1 []string) *students{
	s1 := new(students)
	s1.name = name
	s1.rollno = rollno
	s1.course = course
	s1.array1 = array1
	return s1
}

type studentlist struct{
	list []*students
}

func (ls *studentlist) createstudent(name string, rollno int, course string, array1 []string) *students{
	st2 := NewStudent(name, rollno, course, array1)
	ls.list = append(ls.list, st2)
	return st2
}

func main(){
	slist := new(studentlist)
	slist.createstudent("Huzaifa", 10, "BlockChain", []string{"1"})
	slist.createstudent("Moiz", 20, "CA", []string{"2"})
	fmt.Println(*slist)
}