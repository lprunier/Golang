
package main

import "fmt"

func main() {
	var age int = 0	
	var annee int = 0
	var mois int = 0
	var nom = ""
	var prenom = ""

	for (prenom == "") {
		fmt.Printf("Quel est votre prénom?\n")
		fmt.Scanf("%s", &prenom)
	}
	for (nom == "") {
		fmt.Printf("Quel est votre nom?\n")
		fmt.Scanf("%s", &nom)
	}
	for (annee < 1867) {
		fmt.Printf("Quelle est votre année de naissance?\n")
		fmt.Scanf("%d", &annee)
	}
	for (mois < 1 || mois > 12) {
		fmt.Printf("Quel est votre mois de naissance (1-12)?\n")
		fmt.Scanf("%d", &mois)
	}
	if mois < 6 {
		age = 2017 - annee;
	} else {
		age = 2017 - annee - 1;
	}
	if age < 0 {
		fmt.Printf("%s, vous n'êtes pas né...\n", prenom+" "+nom)
	} else {
		fmt.Printf("%s, vous avez %d ans.\n", prenom+" "+nom, age)
	}
}
