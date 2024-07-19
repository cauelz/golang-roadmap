package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0, errors.New("error reading file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 0, errors.New("error parsing float")
	}

	return value, nil
}

func WriteFloatToFile(value float64, filename string) {
	fmt.Println("Writing Account value to File...")
	valueText := fmt.Sprint(value)
	// What is this file permission?
	// 0644 is a file permission in Unix-like systems. It is a 3-digit octal number.
	// The first digit is the owner's permission, the second digit is the group's permission, and the third digit is the others' permission.
	os.WriteFile("filename.txt", []byte(valueText), 0644)
}