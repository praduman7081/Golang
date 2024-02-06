package main


import "fmt"

// Nameing of gloabal variable start with capital letter and access by anywhaere {Public}
const LoginToken string = "ghdj"

func main(){
     fmt.Println("Variables")

	 // Declerations of variables -->  var name type , var name type = value 
	 var  username string = "Praduman"
	 fmt.Println(username) 
	 fmt.Printf("Variables is type of: %T \n", username )

	 var  isLoggedIn bool = true
	 fmt.Println(isLoggedIn) 
	 fmt.Printf("Variables is type of: %T \n", isLoggedIn )

	 var  smallval int = 256
	 fmt.Println(smallval) 
	 fmt.Printf("Variables is type of: %T \n", smallval )

	 var  fval float64 = 256.5748905746739
	 fmt.Println(fval) 
	 fmt.Printf("Variables is type of: %T \n", fval )

	 // default values and some aliases
	 var  intval int
	 fmt.Println(intval) 
	 fmt.Printf("Variables is type of: %T \n", intval )

	 var  strval string
	 fmt.Println(strval) 
	 fmt.Printf("Variables is type of: %T \n", strval )

	 var  boolval bool
	 fmt.Println(boolval) 
	 fmt.Printf("Variables is type of: %T \n", boolval )

	 // implicit style
	 var website = "LCO"
	 fmt.Println(website)

	 // no var style (block scope or only uses inside a method)
      numberofuser := 200
	  fmt.Println(numberofuser)

	  
}