package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var choix int
	var filename string
	var text string

	for {
		// affichage du menu
		fmt.Println("1 - Récupérer tout le texte contenu dans un fichier")
		fmt.Println("2 - Ajouter du texte dans un fichier")
		fmt.Println("3 - Supprimer tout le contenu d'un fichier")
		fmt.Println("4 - Remplacer le contenu d'un fichier par du texte donné")
		fmt.Println("0 - Quitter")

		// demande de saisie utilisateur
		fmt.Print("Entrez votre choix : ")
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			// récupération du nom de fichier
			fmt.Print("Entrez le nom du fichier : ")
			fmt.Scanln(&filename)

			// lecture du fichier
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier :", err)
				continue
			}

			// affichage du contenu
			fmt.Println(string(data))
		case 2:
			// récupération du nom de fichier
			fmt.Print("Entrez le nom du fichier : ")
			fmt.Scanln(&filename)

			// ouverture du fichier en mode ajout
			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier :", err)
				continue
			}
			defer file.Close()

			// demande de saisie utilisateur
			fmt.Print("Entrez le texte à ajouter : ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				text = scanner.Text()
			}

			// écriture dans le fichier
			_, err = fmt.Fprintln(file, text)
			if err != nil {
				fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			}
		case 3:
			// récupération du nom de fichier
			fmt.Print("Entrez le nom du fichier : ")
			fmt.Scanln(&filename)

			// suppression du fichier
			err := os.Remove(filename)
			if err != nil {
				fmt.Println("Erreur lors de la suppression du fichier :", err)
			} else {
				fmt.Println("Fichier supprimé avec succès.")
			}
		case 4:
			// récupération du nom de fichier
			fmt.Print("Entrez le nom du fichier : ")
			fmt.Scanln(&filename)

			// demande de saisie utilisateur
			fmt.Print("Entrez le nouveau texte : ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				text = scanner.Text()
			}

			// écriture dans le fichier
			err := ioutil.WriteFile(filename, []byte(text), 0644)
			if err != nil {
				fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			} else {
				fmt.Println("Contenu du fichier remplacé avec succès.")
			}
		case 0:
			// sortie de la boucle
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}

		fmt.Println()
	}
}
