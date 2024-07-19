package pointer

import "fmt"

func getAdultYearsByValue(age int) int {
	fmt.Println("----getAdultYearsByValue----")
	fmt.Println("Memory Address inside ", &age)
	return age - 18
}

func getAdultYearsByPointer(age *int) int {
	fmt.Println("----getAdultYearsByPointer----")
	fmt.Println("Memory Address inside ", age)
	*age = *age - 18
	return *age
}