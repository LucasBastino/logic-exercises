package main

import "fmt"

func ejercicio3() {
	// secuencia := make([]int, 50)
	secuencia := [50]int{0, 1}
	for i := 2; i < 50; i++ {
		secuencia[i] = secuencia[i-1] + secuencia[i-2]
	}
	fmt.Println(secuencia)

	// numero0 := 0
	// numero1 := 1

	// for i := 0; i < 50; i++ {
	// 	fmt.Println(numero0)
	// 	sumaTotal := numero0 + numero1
	// 	numero0 = numero1
	// 	numero1 = sumaTotal
	// }
}
