package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Userinput")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(("Enter the rating for our Pizza:"))

	// comma ok // err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T,  ", input)


}