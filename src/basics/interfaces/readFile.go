package interfaces

import "os"

func openFile() {

	data, error := os.ReadFile("texto.txt")

	if error != nil {
		panic(error)
	}

	println(string(data))
}