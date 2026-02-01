package main 

import "fmt"

type Filter func(string) string

func SayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name) 
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// func main () {
// 	SayHelloWithFilter("Eki", spamFilter)
// 	SayHelloWithFilter("anjing", spamFilter)
// } 
// jadi pada kode diatas, function spamFilter di jadikan parameter pada function SayHelloWithFilter
 // dimana function spamFilter akan memfilter kata "anjing"
 // sehingga ketika di panggil pada function SayHelloWithFilter,
 // kata "anjing" akan di ganti menjadi "..."
 // dan hasilnya akan mencetak "Hello ..."
 // sedangkan kata "Eki" tidak akan di ganti karena tidak ada dalam filter
 // sehingga hasilnya akan mencetak "Hello Eki"
 // jadi function bisa di jadikan parameter pada function lain
