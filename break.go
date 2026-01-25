package main

import "fmt"

func Break() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan Ke", i)
	}
}