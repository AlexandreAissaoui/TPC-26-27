package devinette

import (
	"fmt"
	"math/rand"
)

func Devinette(N int) {
	n := rand.Intn(N)
	essai, tries := -1, 1
	for tries <= 5 {
		fmt.Println("Essayez de deviner le nombre entre 1 et", N)
		fmt.Scan(&essai)
		if essai == n {
			fmt.Println("Réussi : ", n)
			break
		}
		tries++
	}
}
