package main

import "fmt"

func main() {
	var nama = "Wawan Qurniawan"
	fmt.Println("Halo ", nama, " selamat belajar golang semangat backend")

	var nilai = [5]int{11, 22, 33, 44, 55}

	for i := 0; i < len(nilai); i++ {
		fmt.Println(nilai[i])
	}

}
