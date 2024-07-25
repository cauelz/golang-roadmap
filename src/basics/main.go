package main

import (
	"basics/arrays"
	"basics/generics"
	"fmt"
)

func main() {

	arrays.Run()
	fmt.Println("----- Comparing Arrays with Slices -----\n\n")
	arrays.ComparingArraysWithSlices()
	generics.Run()
}