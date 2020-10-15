package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Masukan Jumlah Nilai :")
	fmt.Scan(&input)
	data, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Harap Masukan Angka")
		return
	}

	if data%2 == 0 {
		gambarGenap(data)
	} else {
		gambarGanjil(data)
	}
}

func gambarGanjil(data int) {
	tengah := data / 2
	var tmp string
	// fmt.Println("---Panjang---")
	for i := 0; i < data; i++ {
		if i == tengah {
			for k := 0; k < data; k++ {
				tmp += "*"
			}
		} else {
			for w := 0; w < data; w++ {

				if w == 0 {
					tmp += "*"
				} else if w == data-1 {
					tmp += "*"
				} else {
					tmp += "="
				}
			}
		}
		fmt.Println(tmp)
		tmp = ""
	}
}

func gambarGenap(data int) {
	tengah := data / 2
	var tmp string
	// fmt.Println("---Panjang---")
	for i := 0; i < data; i++ {
		if i == tengah || i == tengah-1 {
			for k := 0; k < data; k++ {
				tmp += "*"
			}
		} else {
			for w := 0; w < data; w++ {

				if w == 0 {
					tmp += "*"
				} else if w == data-1 {
					tmp += "*"
				} else {
					tmp += "="
				}
			}
		}
		fmt.Println(tmp)
		tmp = ""
	}
}
