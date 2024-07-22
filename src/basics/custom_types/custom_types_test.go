package custom_types

import "testing"


func TestCustomString(t *testing.T) {
	// Create a new custom string
	var text customString = "Hello, World!"
	
	// Log the custom string
	text.log()
}