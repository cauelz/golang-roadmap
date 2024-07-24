package generics

import "fmt"

func Run() {

	fmt.Println("----- Generics -----")

	PrintAnyValue(1)

	PrintAnyValue("Hello")

	PrintAnyValue(1.0)

	PrintAnyValue(true)
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