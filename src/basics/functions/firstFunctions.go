package functions

// Functions in Go are defined using the func keyword followed by the name of the function.
// The name of the function must be in camel case.
// The parameters of the function are defined inside the parentheses.
// The return type of the function is defined after the parentheses.

// The syntax of a function is as follows:
// func <function_name>(<parameters>) <return_type> {
// 	<code>
// }

// The following function takes two integers as parameters and returns an integer.
func add(x int, y int) int {
	return x + y
}

// There are naked returns in Go, which means that we can define the return variables in the function signature.
// The following function takes two integers as parameters and returns an integer.
func subtract(x int, y int) (result int) {
	result = x - y
	return
}