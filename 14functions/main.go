package main

import "fmt"

func main() {
	fmt.Println("Welcomde in functions in golang")	
	greeter()
   result := adder(3,5)
	fmt.Println("Result is: ", result)

	proRes, message := proAdder(2,3,5,7)
	fmt.Println("Summation is : ", proRes )
	fmt.Println("Pro message is:", message)

	
}

func adder(val1 int, val2 int)(int){
	return val1 + val2
}

func proAdder(values ...int) (int, string){
	total := 0

	for _, val := range values{
		total += val
	}
	return total, "Hii from proadder"
}
func greeter(){
	fmt.Println("Namstey from golang")
}