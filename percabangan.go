package main

import "fmt"

func main() {
	fmt.Println("Materi Perulangan")
	var nilai = 90

	if nilai >= 80 {
		fmt.Println("Kamu Lulus, Karena nilai adalah = ", nilai)
	} else if nilai >= 70 {
		fmt.Println("Kamu Harus Matrikulasi, arena nilai adalah = ", nilai)
		fmt.Println("Kamu pasti bisa,terus berjuang")
	} else {
		fmt.Println("Kamu Tidak Lulus, arena nilai adalah = ", nilai)
	}
}
