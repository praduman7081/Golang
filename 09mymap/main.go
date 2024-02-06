package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	langauges := make(map[string]string)
	langauges["JS"] = "Javascript"
	langauges["PY"] = "Python"
	langauges["RB"] = "Ruby"

	fmt.Println(langauges)
	fmt.Println("JS short for:", langauges["JS"])

	delete(langauges, "PY")
	fmt.Println(langauges)

	// loops

	for key, value := range langauges{
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

}