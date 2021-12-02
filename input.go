package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var letterToRune rune

func input() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Indique moi une letter de ton choix : ")
	scanner.Scan()
	choice := scanner.Text()
	return choice
}

func (p *hangmanData) checkInput() {
	letter := input()
	for _, r := range letter {
		letterToRune = r
	}
	if len(letter) != 1 {
		fmt.Println("Tu peux rentrée un seule caratere à la fois !! Recommence encore une fois l'ami...")
		input()
	} else if !(('a' <= letterToRune && letterToRune <= 'z') || ('A' <= letterToRune && letterToRune <= 'Z')) {
		fmt.Println("Tu peux seulement mettre des carateres alphabetiques !! Recommence encore une fois l'ami...")
		input()
	} else {
		p.CheckLetter(strings.ToLower(letter))
		p.letterlist = append(p.letterlist, strings.ToLower(letter))
	}
}
