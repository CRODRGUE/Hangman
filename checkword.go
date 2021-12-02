package main

import "fmt"

func (p *hangmanData) CheckLetter(letter string) {
	var checkF bool
	for indexFind, letterFind := range p.wordToFind {
		if letter == string(letterFind) {
			checkF = true
			p.word[indexFind] = letter
		}
	}
	if !checkF {
		fmt.Println("La lettre n'est pas presente dans le mot -1 essais")
		p.counter--
	}
}

func (p hangmanData) checkWord() bool {
	var check bool
	for _, letter := range p.word {
		if letter == "_" {
			check = true
		}
	}
	if check && p.counter == 0 {
		fmt.Println("Dommage tu as depasser le nombre maximum d'essais... =(")
		fmt.Println("Le mot à trouver était : ", p.wordToFind)
		check = false
	} else if !check && p.counter != 0 {
		fmt.Println("Bravo l'aim ! tu as trouver le mot sercet...")
		fmt.Println("Qui était : ", p.wordToFind)
	}
	return check
}
