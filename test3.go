package main

import "fmt"

func main() {
	var jam, menit, detik int
	var keterangan string

	fmt.Println("Exp 12:00:00 PM")
	fmt.Println("Masukkan waktu:")
	fmt.Scanf("%d:%d:%d %s", &jam, &menit, &detik, &keterangan)

	if (jam == 12) && ((menit != 00) || (detik != 00)) {
		fmt.Println("Waktu Tidak Valid")
	} else if (jam > 12) || (jam < 0) {
		fmt.Println("Waktu Tidak Valid")
	} else if (menit >= 60) || (jam < 0) {
		fmt.Println("Waktu Tidak Valid")
	} else if (detik >= 60) || (detik < 0) {
		fmt.Println("Waktu Tidak Valid")
	} else if keterangan != "AM" && keterangan != "PM" {
		fmt.Println("Waktu Tidak Valid")
	} else {
		if jam == 12 {
			if keterangan == "AM" {
				fmt.Println("00:00")
			}
			if keterangan == "PM" {
				fmt.Println("12:00")
			}
		} else if menit <= 10 {
			fmt.Printf("%d:0%d\n", jam, menit)
		} else {
			fmt.Printf("%d:%d\n", jam, menit)
		}
	}
	//fmt.Print(keterangan)
}
