package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// inteiros
	fmt.Println(1, 2, 3, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b)) //uint8

	// inteiros com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int64 é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa o mapeamento da tabela unicode int32
	fmt.Println("O tipo de i2 é", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é", i2)

}
