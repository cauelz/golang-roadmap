package custom_types

import "fmt"

type customString string

func (text *customString) log() {
	// Log the custom string
	fmt.Println(*text)
}