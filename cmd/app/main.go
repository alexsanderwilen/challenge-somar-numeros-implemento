package main

import (
	"fmt"
	"github.com/alexsanderwilen/challenge-somar-numeros/internal/somador"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Resultado:", somador.SomarNumeros(nums))
}
