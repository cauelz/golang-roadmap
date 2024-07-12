package arrays

func Array() {
	// Arrays in Go are fixed in size, which means that we need to declare the size of the array
	// when we declare it. We can use the following syntax to declare an array:
	// var <variable_name> [<size>] <type> = [<size>]{<value1>, <value2>, ...}

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

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
}
