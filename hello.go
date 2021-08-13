package main

import (
	"fmt"
)

func main() {
	// tipe data
	// fmt.Println("satu = ", 1)
	// fmt.Println("dua = ", 2)
	// fmt.Println("tiga koma lima = ", 35)

	// tipe data boolean
	// fmt.Println("benar =", true)
	// fmt.Println("benar =", false)

	// string
	// fmt.Println("eko ")
	// fmt.Println("eko kurniawan ")

	// variabel
	// var (
	// 	name1 = "w"
	// nama = true
	// )
	// name1 = "eko"
	// fmt.Println(name1)
	// name = "eko a"
	// fmt.Println(name)
	// fmt.Println(name)
	// fmt.Println(nama)

	// constant
	// const (
	// 	name = 2
	// )
	// fmt.Println(name)

// konversi tipe data
var (
	nilai32 int32 = 257
	nilai64 int64 = int64(nilai32)
	nilai8 int8 = int8(nilai32)

	nilaiuint16 uint16 = uint16(nilai8)
)

fmt.Println(nilai32)
fmt.Println(nilai64)
fmt.Println(nilai8)
fmt.Println(nilaiuint16)


}
