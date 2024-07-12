package mapbasics

import (
	"fmt"
	"testing"
)


func TestMaps(t *testing.T) {

	// What is a map?
	// A map is a collection of key-value pairs. Maps are also known as associative arrays, hash tables, or dictionaries.
	// A map is an unordered collection of key-value pairs. All Keys must be of the same type and all values must be of the same type.
	// but the keys and values can be of different types. The keys must be comparable using ==, so they must be of a type that supports the == operator.

	// How to create a map?
	// var mapName map[keyType]valueType is the syntax to create a map.
	// colors := make(map[string]string) is another way to create a map.
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#00ff00",
	// }

	colors := make(map[string]string)
	// How to access a value in a map?
	// colors["red"] will return the value of the key "red" in the colors map.
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"

	// How to check if a key exists in a map?
	// value, exists := colors["red"] will return the value of the key "red" in the colors map and a boolean value that indicates if the key exists in the map.
	value, exists := colors["red"]
	fmt.Println("--- Accessing a value and checking if a key exists in a map ---")
	fmt.Println("Value: ", value)
	fmt.Println("Exists: ", exists)
	
	fmt.Println("--- printing the colors map ---")
	fmt.Println("Colors Map: ", colors)

	// How to delete a key in a map?
	// delete(colors, "red") will delete the key "red" in the colors map.
	delete(colors, "red")

	// How to iterate over a map?
	// for key, value := range colors { ... } will iterate over the colors map.
	fmt.Println("--- Iterating over a map ---")
	for key, value := range colors {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	// How to iterate over a map and check if a key exists in the map?
	// for key, value := range colors { ... } will iterate over the colors map.
	// if key == "green" { ... } will check if the key is "green".
	fmt.Println("--- Iterating over a map and checking if a key exists in the map ---")
	for key, value := range colors {
		if key == "green" {
			fmt.Println("Key: ", key, "Value: ", value)
		}
	}
	
}

func TestOrderMap(t *testing.T) {

	// How to create a map of a struct?
	// var mapName map[keyType]valueType is the syntax to create a map.
	// orders := make(map[int]Order) is another way to create a map.
	fmt.Println("--- Creating a map of a struct ---")
	orders := make(map[int]Order)
	orders[1] = Order{OrderID: 1, OrderDate: "2021-01-01", OrderAmount: 100.00}
	orders[2] = Order{OrderID: 2, OrderDate: "2021-01-02", OrderAmount: 200.00}
	orders[3] = Order{OrderID: 3, OrderDate: "2021-01-03", OrderAmount: 300.00}
	fmt.Println("Orders: ", orders)
	// How to access a value in a map of a struct?
	fmt.Println("--- Accessing a value in a map of a struct ---")
	fmt.Println("Order 1: ", orders[1])
	fmt.Println("Order 2: ", orders[2])

	// How to iterate over a map of a struct?
	fmt.Println("--- Iterating over a map of a struct ---")
	for key, value := range orders {
		fmt.Println("Key: ", key, "Value: ", value)
	}
}