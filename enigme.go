package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	enigme()
}

func enigme() {
	var reponse []string // création des tableau pour afficher les resultat
	var reponse2 []string
	var reponse3 []string
	var reponse4 []string
	readfile, _ := os.Open("File.txt") // creation d'un tableau ou l'on mettra tout les mots sans les inconveniant du type espace etc...
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string // le tableau
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	for i, c := range lines { // boucle for pour parcourir le tableau
		if i == 0 { // premiere condition pour ajouter le premier element
			reponse = append(reponse, c)
		}
		if i == len(lines)-1 { // deuxieme condition pour ajouter le deuxieme element
			reponse2 = append(reponse2, c)
		}
		if c == "before" { // troisieme condition pour convertir le string apres "before" en int et ajouter le troisieme element en fonction du int donner
			v := lines[i+1]
			a, _ := strconv.Atoi(v)
			reponse3 = append(reponse3, lines[a-1])
		}
		if c == "now " { // quatrieme condition pour renvoyer la 2 eme lettre du mot avant "now " diviser par la totalité du nombre de mot
			b := lines[i-1]
			m := int(b[1]) / len(lines)
			reponse4 = append(reponse4, lines[m-1])
		}

	}
	fmt.Println("Fragment 1:", reponse) // affichage des reponse
	fmt.Println("Fragment 2:", reponse2)
	fmt.Println("Fragment 3:", reponse3)
	fmt.Println("Fragment 4:", reponse4)
	rand.Seed(time.Now().UnixNano()) // creation d'un generateur de nombre random
	fmt.Println(rand.Intn(100))      // limite mis a 100 pour le test
}
