package interfaces

import "testing"

func TestBotInterface(t *testing.T) {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}