package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("linha")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.1415156

	// converte para string
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs) // concatena string com string
	fmt.Println("O valor de x é", x)    //Não precisa converter, pois nao concatena

	fmt.Printf("O valor de x é %.2f", x) // f = float, s = string, d = inteiro, t=boolean, v=generico

	a := 1
	b := 1.99999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
