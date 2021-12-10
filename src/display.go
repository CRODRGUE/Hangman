package main

import (
	"fmt"
	"os"
)

var jose []string

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func (p hangmanData) displayWord() {
	fmt.Println("-------------------")
	fmt.Print("Voici le mot à trouver --> ")
	for _, letter := range p.word {
		fmt.Print(letter, " ")
	}
	fmt.Println("\n")
}

func (p *hangmanData) displayCounter() {
	fmt.Println("tu as encore ", p.counter, " essais")
}

func (p *hangmanData) displayListLetter() {
	fmt.Println()
	fmt.Println("Liste des letters :")
	for i := 0; i < len(p.letterlist)-1; i++ {
		fmt.Print(p.letterlist[i], "/")
	}
	fmt.Println()
	fmt.Println("--------------------")
}

func displayTab(min, max int, Tab []string) {
	for i := min; i <= max; i++ {
		fmt.Println(Tab[i])
	}
}

func (p hangmanData) displayJose() {
	count := p.counter
	switch count {
	case 9:
		displayTab(0, 6, jose)
	case 8:
		displayTab(7, 14, jose)
	case 7:
		displayTab(15, 22, jose)
	case 6:
		displayTab(23, 30, jose)
	case 5:
		displayTab(31, 38, jose)
	case 4:
		displayTab(39, 46, jose)
	case 3:
		displayTab(47, 54, jose)
	case 2:
		displayTab(55, 62, jose)
	case 1:
		displayTab(63, 70, jose)
	case 0:
		displayTab(71, 78, jose)
	}

}
