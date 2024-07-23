package main

import (
	"bufio"
	"fmt"
	"notepad_project/note"
	"os"
	"strings"
)

func main() {
	title, body := getNoteData()

	note, err := note.New(title, body)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note saved successfully.")
}

func getNoteData() (string, string) {
	title := ReadDataFromUser("Enter the title of the note: ")

	body := ReadDataFromUser("Enter the body of the note: ")

	return title, body
}

func ReadDataFromUser(prompt string) string {

	fmt.Print(prompt)

	// Read data from the user input.
	// bufio.NewReader returns a new Reader whose buffer has the default size, which means it reads from the standard input.
	// os.Stdin is a File type which is an open file descriptor for standard input, which means it reads from the standard input.
	reader := bufio.NewReader(os.Stdin)

	// ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter.
	data, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.
	// TrimSuffix(data, "\n") removes the newline character from the data.
	// TrimSuffix(data, "\r") removes the carriage return character from the data.
	text := strings.TrimSuffix(data, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}
