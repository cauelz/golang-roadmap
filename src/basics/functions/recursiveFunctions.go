package functions

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

	// Factorial with for loop
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result *= i
	// }
}