package main

// import "fmt"

// func endApp() {
// 	fmt.Println("End App")
// }

// func runApp(error bool) {
// 	defer endApp()

// 	if error {
// 		panic("Aplikasi Error")
// 	}
// }

// // func main() {
// // 	runApp(true)
// // }
// // jadi ketika panic terjadi, program akan langsung berhenti dan mengeksekusi deferred function sebelum benar-benar berhenti.
// // sehingga outputnya adalah:
// // End App
// // panic: Aplikasi Error
// // jadi function endApp tetap dieksekusi meskipun terjadi panic pada function runApp.
// // ini berguna untuk membersihkan resource atau melakukan tindakan tertentu sebelum program benar-benar berhenti akibat panic.
// // jadi defer dan panic bisa digunakan bersama-sama untuk memastikan bahwa tindakan tertentu selalu dilakukan sebelum program berhenti akibat panic.
// // panic adalah mekanisme untuk menghentikan eksekusi program secara tiba-tiba ketika terjadi kondisi yang tidak diinginkan atau error yang serius.
// // ketika panic terjadi, program akan berhenti dan menampilkan pesan error yang disebabkan oleh panic tersebut.