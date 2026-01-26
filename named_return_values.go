package main

import "fmt"

func getCompleteName(firstName, middleName, lastName string) (string, string, string) {
	firstName = "Eki"
	middleName = "Adi"
	lastName = "Ramadhan"

	return firstName, middleName, lastName
}

// func main() {
// 	a, b, c := getCompleteName("", "", "")
// 	fmt.Println(a, b, c)
// }