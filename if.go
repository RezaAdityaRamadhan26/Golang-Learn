package main

import "fmt"

func IfElse() {
	name:= "Eki"

	if name == "Eki" {
		fmt.Println("Hello Eki")
	} // if statement tanpa else

	if name == "Eki" {
		fmt.Println("Hello Eki")
	} else {
		fmt.Println("Hello, Orang lain")
	} // if statement dengan else

	if name == "Eki" {
		fmt.Println("Hello Eki")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} // if statement dengan else

	


}