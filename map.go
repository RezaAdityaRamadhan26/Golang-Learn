package main

import "fmt"

func Mapping() { 
	// ======= Kode Map Normal =======
	// person := map[string]string{
	// 	"name":    "Eko",
	// 	"address": "Jawa Barat",
	// } // Deklarasi map kode lebih banyak

	// orang := map[string]string {
	// 	"nama": "Eko",
	// 	"alamat": "Jawa Barat",
	// } // Deklarasi map kode lebih sedikit

	// fmt.Println(person)
	// fmt.Println(orang)

	book := make(map[string]string)

	book["title"] = "Belajar Gila"
	book["author"] = "Agung"
	book["ups"] = "salah"

	fmt.Println(book)

	// ======= Kode Map Dengan Delete =======
	delete(book, "ups")
	fmt.Println(book)

	// ======= Kode Map Dengan Looping =======
	for key, value := range book {
		fmt.Println(key, ":", value)
	}

	// ======= Kode Map Dengan Mengecek Key =======
	author, isExist := book["author"]
	if isExist {
		fmt.Println("Author:", author)
	} else {
		fmt.Println("Author tidak ditemukan gung")
	}
}

