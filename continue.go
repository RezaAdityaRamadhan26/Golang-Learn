package main

import "fmt"

func Continue() {
	for i := 0; i < 10; i++  { // Increment i++
		if i%2 == 0 {
			continue // Skip 2 angka setiap perulangan (1, 3, 5, DSB)
		}
		fmt.Println("Perulangan Ke", i)
	}
}

// func main() {
// 	fmt.Println(Continue)

// 	Continue()
// }
