package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}


func main() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 // maka hasilnya akan sama dengan loop diatas
	fmt.Println(result)
	fmt.Println(factorialLoop(10)) // hasilnya adalah perkalian dibawah 10
								   // 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 = 3628800
	fmt.Println(factorialRecursive(10)) // jadi hasilnya sama dengan loop diatas, tetapi menggunakan recursive
										// recursive adalah function yang memanggil dirinya sendiri
										// pada contoh diatas, function factorialRecursive memanggil dirinya sendiri
										// hingga mencapai nilai dasar yaitu 1
										// sehingga hasilnya sama dengan loop diatas
}