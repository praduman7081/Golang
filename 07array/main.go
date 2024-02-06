package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Gavava"
	fruitList[3] = "Mango"
	fruitList[4] = "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
}