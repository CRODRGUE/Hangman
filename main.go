package main

import "fmt"

type hangmanData struct {
	wordToFind string
	word       []string
	letterlist []string
	counter    int
}

func main() {
	var player hangmanData
	player.init()
	player.showRandLetter()
	initJose()
	fmt.Println("===================================================")
	fmt.Println("---> ", player.wordToFind)
	fmt.Println(player.word)
	fmt.Println("===================================================")

	for player.checkWord() {
		player.displayWord()
		player.displayListLetter()
		player.displayCounter()
		player.displayJose()
		player.checkInput()
		//Clear()
	}
}
