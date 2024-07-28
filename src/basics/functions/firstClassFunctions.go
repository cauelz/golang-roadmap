package functions

type TransformerFn func(int) int

func TransformNumbers(numbers *[]int, transformerFn TransformerFn) []int {

	var transformedNumbers []int

	//var transformerFn = getTransformerFn(transformer)

	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, transformerFn(number))
	}

	return transformedNumbers
}

func CreateTransformerFn(factor int) TransformerFn {
	return func(number int) int {
		return number * factor
	}
}

// func getTransformerFn(transformer string) TransformerFn {
// 	switch transformer {
// 	case "double":
// 		return Double
// 	case "triple":
// 		return Triple
// 	default:
// 		return nil
// 	}
// }

// func Double(number int) int {
// 	return number * 2
// }

// func Triple(number int) int {
// 	return number * 3
// }