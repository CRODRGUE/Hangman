package main

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
	/* Pour faire des tests (affiche le mot au lancement !)
	fmt.Println("===================================================")
	fmt.Println("---> ", player.wordToFind)
	fmt.Println(player.word)
	fmt.Println("===================================================") */

	for player.checkWord() && player.counter > 0 {
		player.displayWord()
		player.displayListLetter() // Les lettres s'affiche avec un tour de retard mais impossible de trouver pourquoi ?
		player.displayCounter()    // Si tu trouves le probleme dis le moi =)
		player.displayJose()       // (J'ai regarder 10 fois comment s'execute le programme normalement cela ne vient de sa position !)
		player.checkInput()
	}
	player.displayJose()
}
