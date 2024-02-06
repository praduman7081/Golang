package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang, No super or parent

	ankit := User{"Ankit", "gandu@mail.com", false, 22}

	fmt.Println(ankit)
	fmt.Printf("Ankit details are: %+v\n", ankit)
	fmt.Printf("Name is %v and email is %v", ankit.Name, ankit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
