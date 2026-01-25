package main

import "fmt"

func Slicing() {
	// ======= Kode Slice Normal =======
	nama := []string{"yuli", "budi", "firman", "eki", "dani", "agung"}

	slice1 := nama[2:4] // mengambil elemen dari index ke 2 sampai index ke 4 (tidak termasuk index ke 4)

	slice2 := nama[2:] // mengambil elemen dari index ke 2 sampai akhir

	slice3 := nama[:4] // mengambil elemen dari awal sampai index ke 4 (tidak termasuk index ke 4)

	slice4 := nama[:] // mengambil semua elemen

	var slice5 []string = nama[1:3] // deklarasi slice dengan tipe data

	fmt.Println(slice1, slice2, slice3, slice4, slice5)

	// ======= Kode Slice Dengan Append =======

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[2] = "Libur Lagi"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	// perubahan pada slice2 tidak mempengaruhi array days karena terjadi alokasi ulang memory pada saat append 
	// Isi array tidak dapat berubah karna slice

	// ======= Kode Slice Dengan Make =======
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Agung"
	newSlice[1] = "Agung"
	// newSlice[2] = "Eko" // tidak bisa karna harus menggunakan append untuk menambah elemen di luar kapasitas

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eki")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// ======= Kode Slice Copy Slice =======
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// ======= Kesalahan ketika menggunakan Slice, jangan malah membuat array baru =======
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5} // ini salah, karena membuat array baru bukan slice

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}