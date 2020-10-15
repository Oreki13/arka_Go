package main

import (
	"fmt"
	"log"
	"strconv"
)

type create interface {
	genKanan() string
	getKiri() string
}

type rataKanan struct {
	bintang string
	param   int
	length  int
}

type rataKiri struct {
	bintang string
	param   int
}

func (r *rataKanan) genKanan() string {
	for b := 1; b < r.length+1; b++ {
		if b < r.param {
			r.bintang += " "
		} else {
			r.bintang += "*"
		}
	}
	return r.bintang
}

func (r *rataKiri) getKiri() string {
	for i := 0; i < r.param; i++ {
		r.bintang += "*"
	}
	return r.bintang
}

func main() {

	var input string
	fmt.Println("Masukan Angka")
	fmt.Scan(&input)
	panjang, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Harap Masukan Angka")
		return
	}

	var k rataKiri
	var l rataKanan

	for w := 0; w < panjang; w++ {
		k.param = w + 1

		log.Println(k.getKiri())

		k.bintang = ""
	}
	l.length = panjang
	for w := 0; w < l.length; w++ {

		l.param = w + 1
		log.Println(l.genKanan())

		l.bintang = ""
	}

}
