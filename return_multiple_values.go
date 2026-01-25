package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	// return "Reza", "Aditya"
	firstName = "Muhammad"
	middleName = "Reza"
	lastName = "Aditya"

	return firstName, middleName, lastName
}

func main() {
	// firstName, lastName := getFullName() // Memanggil 2 variable yang di return pada function getFullName
	// fmt.Println(firstName, lastName)

	// firstName, _ := getFullName()
	// fmt.Println(firstName) // Memanggil hanya 1 variable, karna tidak bisa 1 diganti dengan "_"

	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}