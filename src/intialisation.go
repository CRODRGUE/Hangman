package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func (p *hangmanData) init() {
	var words []string
	fichier, err := os.Open("../doc/words.txt")
	scanner := bufio.NewScanner(fichier)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	fichier.Close()
	nbr := randNumber((len(words) - 1))
	p.wordToFind = string(words[nbr])
	for i := 0; i < len(p.wordToFind); i++ {
		p.word = append(p.word, "_")
	}
	p.counter = 10
}

func (p *hangmanData) showRandLetter() {
	n := (len(p.wordToFind)/2 - 1)
	for i := 0; i < n; i++ {
		index := randNumber(len(p.wordToFind) - 1)
		p.word[index] = string(p.wordToFind[index])
	}
}

func initJose() {
	fichier, err := os.Open("../doc/hangman.txt")
	scanner := bufio.NewScanner(fichier)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		jose = append(jose, scanner.Text())
	}
	fichier.Close()
}

func randNumber(max int) int {
	time.Sleep(3)
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-0+1) + 0)
}
