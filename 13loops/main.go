package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
    
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	fmt.Println(days)

	// for d:= 0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days{
	// 	fmt.Printf("index is  and value is %v\n",  day)
	// }

	var val int = 1
    // Act like while loop
	for val <10{
		fmt.Println(val)
		val++
	}

}