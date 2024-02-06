package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang, No super or parent

	ankit := User{"Ankit", "gandu@mail.com", false, 22}

	fmt.Println(ankit)
	fmt.Printf("Ankit details are: %+v\n", ankit)
	fmt.Printf("Name is %v and email is %v", ankit.Name, ankit.Email)
	ankit.GetStatus()
	ankit.NewMail()
	fmt.Printf("Name is %v and email is %v", ankit.Name, ankit.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
   fmt.Println("Is user active :", u.Status)
}
func (u User) NewMail(){
    u.Email = "ankit@mail.dev"
	fmt.Println("Email of this user is:", u.Email)
}