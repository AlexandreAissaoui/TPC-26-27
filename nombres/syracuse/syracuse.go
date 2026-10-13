package syracuse

import "fmt"

func Syracuse(N int) (t []int, duree int, altitude int) {
	t, duree, altitude = make([]int, 0, 100), 1, N
	s := N
	t = append(t, s)
	for s != 1 {
		if s%2 == 0 {
			s = s / 2
		} else {
			s = 3*s + 1
		}
		if s > altitude {
			altitude = s
		}
		t = append(t, s)
		duree++
	}
	t = t[0:duree]
	return
}

func CalculeVol(n int) (int, int) {
	_, d, a := Syracuse(n)
	return d, a
}

func AfficheVol(n int) {
	d, a := CalculeVol(n)
	fmt.Println("duree = ", d, " altitude = ", a)
}

func Syracuse2(N int, K int) ([]int, int, int, int) {
	next := func(x int) int {
		if x%2 == 0 {
			return x / 2
		}
		return 3*x + K
	}

	t, altitude := []int{N}, N
	s := N
	duree := 0
	startindex := 0
	f := 0

	for s != 1 {
		s = next(s)
		t = append(t, s)
		duree++

		// Suivi de l'altitude maximale en continu
		if s > altitude {
			altitude = s
		}
		if f == 0 {
			if t[startindex] == t[duree] {
				// Cycle détecté: calcul de sa longueur
				cycle, e, lencycle := []int{t[duree]}, next(t[duree]), 1
				for t[startindex] != e {
					cycle = append(cycle, e)
					e = next(e)
					lencycle++
				}
				// Recherche de la première occurrence de la valeur pour trouver le début du cycle
				foundIndex := -1
				for i, val := range t {
					for j := 0; j < len(cycle); j++ {
						if val == cycle[j] {
							foundIndex = i
							break
						}
					}
					if foundIndex != -1 {
						break
					}
				}

				// On récupère ainsi la durée de la sous-suite de syracuse et de son premier cycle
				duree = foundIndex + lencycle + 1
				t = t[:duree]
				return t, duree, lencycle, altitude
			}
			startindex++
			f = 1
		} else {
			f = 0
		}
	}
	return t, duree + 1, 0, altitude
}

func Syracuse3(N int, K int) ([]int, int, int, int) {
	// Utilisation d'une map pour améliorer l'efficacité de la fonction précédente
	return make([]int, 0), 0, 0, 0
}
