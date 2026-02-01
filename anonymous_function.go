package main

import (
	"fmt"
	_ "fmt"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) { 
	if blacklist(name) { 
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func _main() {

	Blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("eko", Blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
// jadi pada kode diatas, kita membuat function anonymous yang langsung di passing ke parameter function registerUser
// function anonymous tersebut berfungsi sebagai filter untuk memblokir user dengan nama "anjing"
// sehingga ketika kita memanggil registerUser dengan nama "eko", maka akan mencetak "Welcome eko"
// sedangkan ketika kita memanggil registerUser dengan nama "anjing", maka akan mencetak "You are blocked anjing"
// jadi function anonymous ini berguna ketika kita hanya membutuhkan function sekali pakai saja tanpa harus mendefinisikan function secara terpisah

// function anonymous adalah function yang tidak memiliki nama
// function anonymous biasanya digunakan sebagai parameter pada function lain atau sebagai return value pada function
// function anonymous juga bisa disimpan dalam variable dan dipanggil seperti function biasa

// contoh function anonymous yang disimpan dalam variable
var sayHello = func(name string) {
	fmt.Println("Hello", name)
}