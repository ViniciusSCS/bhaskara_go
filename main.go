package main

import (
	"bhaskara/bhaskara"
	"fmt"
)

func main() {
	// Exemplo de valores para os coeficientes a, b e c
	a := 1.0
	b := 8.0
	c := -9.0

	// Calcular as raízes
	x1, x2, sucesso := bhaskara.Bhaskara(a, b, c)

	if sucesso {
		fmt.Printf("As raízes são x1 = %.2f e x2 = %.2f\n", x1, x2)
	} else {
		fmt.Println("Não há raízes reais.")
	}
}
