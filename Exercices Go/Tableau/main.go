package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// fonction qui génère un tableau à n dimensions rempli de valeurs aléatoires
func generateTableau(n int) [][]int {
	if n < 3 {
		n = 3
	}
	if n > 5 {
		n = 5
	}

	// création du tableau
	tableau := make([][]int, n)
	for i := 0; i < n; i++ {
		tableau[i] = make([]int, n)
	}

	// remplissage du tableau avec des valeurs aléatoires
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tableau[i][j] = rand.Intn(100)
		}
	}

	return tableau
}

// fonction qui crée le code HTML pour afficher le tableau dans une page web
func generateHTMLTable(tableau [][]int) string {
	var sb strings.Builder

	sb.WriteString("<table border=\"1\">")
	for i := 0; i < len(tableau); i++ {
		sb.WriteString("<tr>")
		for j := 0; j < len(tableau[i]); j++ {
			sb.WriteString("<td>")
			sb.WriteString(strconv.Itoa(tableau[i][j]))
			sb.WriteString("</td>")
		}
		sb.WriteString("</tr>")
	}
	sb.WriteString("</table>")

	return sb.String()
}

func main() {
	// génération du tableau
	n := rand.Intn(3) + 3
	tableau := generateTableau(n)

	// création du code HTML
	html := generateHTMLTable(tableau)

	// affichage dans une page web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html)
	})
	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
