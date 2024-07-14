package channels

import (
	"fmt"
	"net/http"
)


var links = []string{
	"http://www.google.com", 
	"http://wwww.amazon.com", 
	"http://www.facebook.com",
	"http://www.golang.org",
}

func checkWebsites() {

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// What this code does? It will wait for the channel to return a value and then assign it to l
	// and then call the checkLink function with the value of l
	// This is a blocking call and the main function will wait for the channel to return a value
	// the sintax <-c is a blocking call
	for l := range c{
		go func(l string, c chan string) {
			checkLink(l, c)
		}(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down!")
		c <- link
	}

	fmt.Println(link, "is up!")
	c <- link
}