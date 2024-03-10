package main

import (
	"fmt"
)

func ejercicio1() {
	fmt.Println("Hola mundo!")
	for i := 1; i-1 < 100; i++ {
		if i%15 == 0 {
			fmt.Println("multiplo de 3 y 5")
		} else if i%3 == 0 {
			fmt.Println("multiplo de 3")
		} else if i%5 == 0 {
			fmt.Println("multiplo de 5")
		} else {
			fmt.Println(i)
		}

	}
}
