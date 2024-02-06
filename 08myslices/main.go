package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to study on slices")

	var fruitList = []string{"Apple", "Banana", "Papaya"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	fruitList = append(fruitList, "Tomato", "Peach")
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList[1:4] )
	fmt.Println(fruitList)

	highscore := make([]int, 4)
	highscore[0] = 60
	highscore[1] = 43
	highscore[2] = 57
	highscore[3] = 40
	// highscore[4] = 70

	highscore = append(highscore, 68, 70, 56)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))
    sort.Ints(highscore)
    
	fmt.Println(highscore)

	// how to remove value from slices based on index

	var courses = []string{"reactjs", "java", "ruby", "python"}

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)



}