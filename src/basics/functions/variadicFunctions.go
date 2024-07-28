package functions


func SumValues(initialValue int, values ...int) int {
	sum := initialValue
	for _, value := range values {
		sum += value
	}
	return sum
}