package main

import (
	"fmt"
)

func main() {
	var angka1, angka2 int

	fmt.Print("Masukan angka pertama: ")
	fmt.Scan(&angka1)

	fmt.Print("Masukan angka kedua: ")
	fmt.Scan(&angka2)

	hasil := Tambah(angka1, angka2)
	fmt.Printf("Hasil Penjumlahan: %d\n", hasil)

}
