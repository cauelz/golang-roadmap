package arrays

import "fmt"


func RunExercise1() {
	
	// 1) Create an Array with 3 hobbies of your choice. Print the array.
	hobbies := [3]string{"reading", "training", "coding"}
	fmt.Println("Hobbies Array: ", hobbies)

	// 2) Output the first and last element of the array.
	newHobbiesList := []string{hobbies[0], hobbies[2]}
	fmt.Println("First and Last Hobbies: ", newHobbiesList)

	// 3) Output the last 2 elements of the array in a new list
	lastTwoHobbies := hobbies[1:]
	firstTwoHobbies := hobbies[:2]
	fmt.Println("Last Two Hobbies: ", lastTwoHobbies)
	fmt.Println("First Two Hobbies: ", firstTwoHobbies)
	fmt.Println("Original Hobbies Array: ", hobbies)

	lastTwoHobbies2 := append([]string{}, hobbies[1:]...)
	fmt.Println("Last Two Hobbies 2: ", lastTwoHobbies2)
	fmt.Println("Original Hobbies Array: ", hobbies)

	// 4) Slicing the lasTwoHobbies slice, output the elements.
	lastTwoHobbiesSlice := lastTwoHobbies[0:]
	fmt.Println("Last Two Hobbies Slice: ", lastTwoHobbiesSlice)
	lastTwoHobbiesSlice2 := firstTwoHobbies[1:3]
	fmt.Println("Last Two Hobbies Slice 2: ", lastTwoHobbiesSlice2)

	// 5) Create a Dynamic Array that contains my 2 goals with this course.
	goals := []string{"learn go", "build a project"}
	fmt.Println("Goals: ", goals)
	firstGoal := []string{goals[0]}
	fmt.Println("First Goal: ", firstGoal)
	goals = append(goals, "get a job")
	fmt.Println("Goals: ", goals)



}