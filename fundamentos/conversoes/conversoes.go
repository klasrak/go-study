package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // convertendo int para float64

	fmt.Println("teste" + strconv.Itoa(123)) // int para string

	num, _ := strconv.Atoi("123") // string para int

	fmt.Println(num)
}
