package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entrez le nombre d'allumettes initial : ")
	fmt.Scan(&n)

	if n < 4 {
		fmt.Println("Le nombre d'allumettes doit être au moins 4.")
		return
	}

	var player int
	fmt.Print("Qui commence ? (1 pour vous, 2 pour l'ordinateur) : ")
	fmt.Scan(&player)

	for n > 0 {
		fmt.Printf("\nIl reste %d allumettes.\n", n)

		if player == 1 {
			var take int
			for {
				fmt.Print("Combien d'allumettes prenez-vous ? (1 à 3) : ")
				fmt.Scan(&take)
				if take >= 1 && take <= 3 && take <= n {
					break
				}
				fmt.Println("Mauvaise réponse, entrez un nombre entre 1 et 3 inclus.")
			}
			n -= take
			fmt.Printf("Vous avez pris %d allumettes.\n", take)
			player = 2
		} else {
			take := ((n - 1) % 4) + 1
			if take > n {
				take = n
			}
			n -= take
			fmt.Printf("L'ordinateur a pris %d allumettes.\n", take)
			player = 1
		}
	}

	if player == 1 {
		fmt.Println("\nVous avez gagné !")
	} else {
		fmt.Println("\nL'ordinateur a gagné.")
	}
}
