package main
import "fmt"

func main(){
	// Task # 04
	fmt.Println("Task # 04");
	numbers := []int{999,99,9}
	fmt.Println("Printing Numbers array with 1 increment")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i] + 1)
		fmt.Printf("")
	}
	fmt.Println("... Array Done");
	// Task # 05
	fmt.Println("Task # 05");
	result := true
	name := "Spark"
	size := 2000
	fmt.Println("")
	fmt.Printf("Result = %v, Name = %v, Size = %v", result, name, size)
	// Task # 06
	fmt.Println("Task # 06");
	var array1 = [...]int{10, 20, 30}
	array2 := [5]int{1,2,3,4,5}
	var array3 = [5]string{"huzaifa", "abdullah", "ubaid", "sameel","musab"}
	fmt.Println("Array1: ",array1);
	fmt.Println("Array2: ",array2);
	fmt.Println("Array3: ",array3);
	fmt.Println("0th of array3: ",array3[0]);
	fmt.Println("2nd of array2: ",array2[2]);
	// Task # 07
	fmt.Println("")
	strArray1 := [3]string{"Ali", "Ahmad", "baqir"}
	strArray2 := strArray1
	strArray3 := &strArray1
	fmt.Println("Array1: ", strArray1)
	fmt.Println("Array2: ", strArray2)
	strArray1[0] = "Huzaifa"
	fmt.Println("Array1 after change: ", strArray1)
	fmt.Println("Array2 after change: ", strArray2)
	fmt.Println("Array3 after change: ", strArray3)
}