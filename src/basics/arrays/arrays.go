package arrays

import "fmt"

func Run() {
	// Arrays in Go are fixed in size, which means that we need to declare the size of the array
	// when we declare it. We can use the following syntax to declare an array:
	// var <variable_name> [<size>] <type> = [<size>]{<value1>, <value2>, ...}

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers: ",numbers)

	// It is possible to slice an Array
	// The syntax is <array>[<start>:<end>]
	// A few variations of the slicing syntax:
	// <array>[:<end>] - slice from the beginning to the end
	// <array>[<start>:] - slice from the start to the end
	numberSliced := numbers[1:3]
	fmt.Println("Numbers Sliced from 1 to 3", numberSliced)

	// Web can iterate over an array using a for loop
	// the sintax is the same as in other languages like C, C++, Java, etc.
	// for <index>, <value> := range <array> {
	// 	<code>
	// }
	// The range keyword returns the index and the value of the array at that index
	// We can use the _ character to ignore the index or the value

	// Why do we user the short variable declaration syntax in a for loop?
	// The short variable declaration syntax is used to declare a variable and assign a value to it
	// in the same line. This is useful when we want to declare a variable that is only used in the
	// for loop. This is a common pattern in Go and it's used to keep the code clean and concise.
	
	// Every iteration creates a new variable i and number, so we don't need to declare them outside of the loop
	for i, number := range numbers {
		println(i, number)
	}

	// The example above shows us that an Array is a fixed-size collection of elements of any type.
	// This is possible due because we are using the "any" type, which is a special type in Go that can
	// hold any value. This is useful when we want to create a collection of elements of different types.
	var anyTypeArray [5]any = [5]any{1, "two", 3.0, true, []int{1, 2, 3}}

	for _, value := range anyTypeArray {
		println(value)
	}

	// We can also create a multi-dimensional array in Go. A multi-dimensional array is an array of arrays.
	// We can use the following syntax to declare a multi-dimensional array:
	// var <variable_name> [<size1>][<size2>]...<type> = [<size1>][<size2>]{...}
	var multiDimArray [2][2]int = [2][2]int{{1, 2}, {3, 4}}

	for i, row := range multiDimArray {
		for j, number := range row {
			println(i, j, number)
		}
	}

	// We can also create an array of structs in Go. A struct is a user-defined type that represents a collection
	// of fields. We can use the following syntax to declare an array of structs:
	// var <variable_name> [<size>]<struct_name> = [<size>]<struct_name>{<struct1>, <struct2>, ...}
	type Person struct {
		name string
		age  int
	}

	var people [2]Person = [2]Person{{"Alice", 25}, {"Bob", 30}}

	for _, person := range people {
		println(person.name, person.age)
	}
}


func ComparingArraysWithSlices() {

	//Let's first declare an array of integers
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Numbers Array: ", numbers)
	fmt.Println("Length of Numbers Array: ", len(numbers))
	fmt.Println("Capacity of Numbers Arrays: ", cap(numbers))

	//Now let's slice the array from index 1 to the end
	numberSliced := numbers[1:]

	fmt.Println("Numbers Sliced from 1 to the end: ", numberSliced)
	fmt.Println("Length of numberSliced: ", len(numberSliced))
	fmt.Println("Capacity of numberSliced: ", cap(numberSliced))

	//Can you notice the difference between the length and capacity of the array and the slice?
	//The length of the array is the number of elements in the array
	//The capacity of the array is the total number of elements in the array

	//The length of the slice is the number of elements in the slice
	//The capacity of the slice is the total number of elements in the array starting from the index where the slice starts to the end of the array

	//But what happens if we slice the slices numberSliced from the beginning to the index 1?
	numberSliced = numberSliced[:1]

	fmt.Println("Numbers Sliced from the beginning to 1: ", numberSliced)
	fmt.Println("Length of numberSliced: ", len(numberSliced))
	fmt.Println("Capacity of numberSliced: ", cap(numberSliced))

	//The length of the slice is the number of elements in the slice
	//The capacity of the slice is the total number of elements in the numberSliced starting from the index where the slice starts to the end of the array
	//This happens because the slice is a reference to the array, so when we slice the slice, we are slicing the array. 
	//And the capacity of the slice is the total number of elements in the array starting from the index where the slice starts to the end of the array


}
