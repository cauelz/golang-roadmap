package generics

import "fmt"

func Run() {

	fmt.Println("----- Generics -----")

	PrintAnyValue(1)

	PrintAnyValue("Hello")

	PrintAnyValue(1.0)

	PrintAnyValue(true)

	PrintAnyValue(add(1, 2))

	PrintAnyValue(add(1.0, 2.0))

	PrintAnyValue(add("Hello", " World"))
}
// add is a generic function that accepts any type that is an integer, float64, or string
// and returns the sum of the two values.
// A brief explanation of how it works: 
// The function signature is add[T int | float64 | string](a, b T) T
// The [T int | float64 | string] part is the type constraint. 
// It tells the compiler that the function can only accept types that are either int, float64, or string.
func add[T int | float64 | string](a, b T) T {
	return a + b
}

// PrintAnyValue prints the value of any type because it accepts an interface{}
func PrintAnyValue(value interface{}) {
	
	// Type assertion to check if the value is an integer
	// If it is, print it
	typedValueInt, isInt := value.(int)
	typedValueStr, isStr := value.(string)
	typedValueFloat, isFloat := value.(float64)
	typedValueBool, isBool := value.(bool)

	if isInt {
		println("Print TypedValue: ", typedValueInt, "\n\n")
		return
	}

	if isStr {
		println("Print TypedValue: ", typedValueStr, "\n\n")
		return
	}

	if isFloat {
		println("Print TypedValue: ", typedValueFloat, "\n\n")
		return
	}

	if isBool {
		println("Print TypedValue: ", typedValueBool, "\n\n")
		return
	}
	
}