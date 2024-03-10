package main

import (
	"fmt"
	"strings"
)

func ejercicio2() {
	fmt.Println(sonAnagramas("serugiran", "siganeris"))
}

func sonAnagramas(palabra1, palabra2 string) bool {
	if len(palabra1) != len(palabra2) {
		return false
	} else {
		for i := 0; i < len(palabra1); i++ {
			if strings.Contains(palabra2, string(palabra1[i])) {
				palabra2 = strings.Replace(palabra2, string(palabra1[i]), "", 1)
				continue
			} else {
				return false
			}
		}
		return true
	}
}
