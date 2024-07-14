package channels

import (
	"testing"
)


func TestHttpCalls(t *testing.T) {
	// The following code will fail because the function CheckWebsites() is not a blocking function
	checkWebsites()
	// The following code will pass because the
}