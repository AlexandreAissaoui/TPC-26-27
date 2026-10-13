package main

import (
	"fmt"
	"nombres/syracuse"
)

func main() {
	for i := 90; i < 93; i++ {
		for j := 3; j <= 11; j += 2 {
			t, d, lencycle, _ := syracuse.Syracuse2(i, j)
			fmt.Println(t)
			fmt.Println("duree de la sous-suite de Syracuse, terminant par un cycle :", d)
			fmt.Println("longueur du cycle =", lencycle)
		}
	}
	//devinette.Devinette(10)
}
