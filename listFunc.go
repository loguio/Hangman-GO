package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func PrintHangman(fileName string, attempts int) { // fonction qui ecrit dans la console le pendu a chaque fois que l'on perd une vie
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("could not find any file")
	}
	nbrattemps := 10 - attempts
	tabword := make([]byte, 80)
	file.Seek(int64(80*nbrattemps), 0)
	file.Read(tabword)
	hangman := string(tabword)
	fmt.Println(hangman)
}

func LenNmbtoPrint(lenWord int) int {
	return lenWord/2 - 1
}

func randomint(word string) int { // fonction qui sors un chiffre aléatoirement sur le maximum de lettre à deviné
	var randomint int
	var max = len(word)
	rand.Seed(time.Now().UnixNano())
	randomint = int(rand.Int31n(int32(max)))
	return randomint
}

func getUserletter() []string { // fonction qui demande au joueur une lettre
	fmt.Println("\nEnter a valid letter")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}
	input = strings.TrimSuffix(input, "\n")
	tabstr := strings.Split(input, "")
	return tabstr
}

func tabtostr(tab []string) string { // convertie un tableau de string en string simple
	word := ""
	for i := 0; i < len(tab); i++ {
		word += tab[i]
	}
	return word
}

func GetRandomWord(fileName string) string { // fonction qui prend un mot aléatoire dans la list de mot
	file, err := os.Open(fileName)
	if err != nil {
		return "could not find any file"
	}
	rand.Seed(time.Now().UnixNano())
	var i int64
	nbNl := 0
	nbchar := 0
	testchar := make([]byte, 1)
	var word string
	indexword := 0
	for indexword == 0 {
		indexword = int(rand.Int31n(38))
	}
	index := 0
	for i = 0; nbNl < int(indexword); i++ {
		file.Seek(i, 0)
		file.Read(testchar)
		strchar := string(testchar)
		if strchar == "\n" {
			nbNl++
			if nbNl != int(indexword) {
				nbchar = 0
				index = int(i + 1)
			}
		} else {
			nbchar++
		}
	}
	tabword := make([]byte, nbchar)
	file.Seek(int64(index), 0)
	file.Read(tabword)
	word = string(tabword)
	return word
}
