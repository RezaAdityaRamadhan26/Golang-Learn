package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	} 
	// message := recover() // ini adalah recover yang salah
	// fmt.Println("terjadi panic", message)
}

func main() {
	runApp(true)
}