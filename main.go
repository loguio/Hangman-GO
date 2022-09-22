package main

import (
	"fmt"
	"strings"
)

func main() {
	attemtps := 10
	fmt.Printf("Good luck you have %v attemtps \n", attemtps)
	word := GetRandomWord("wordsCopy.txt")
	wordtab := strings.Split(word, "")
	fmt.Println(len(wordtab))
	LenNmbtoPrint := LenNmbtoPrint(len(word))

	var emptyword []string // tableau de rune vide pour le pendu
	for i := 0; i < len(wordtab)-1; i++ {
		emptyword = append(emptyword, "_")
		emptyword = append(emptyword, " ")
	}
	emptyword = append(emptyword, "_")

	for i := 0; i < LenNmbtoPrint; i++ {
		nmrandom := randomint(word)
		k := nmrandom * 2
		emptyword[k] = string(wordtab[nmrandom])
	}
	find := 0
	emptywordfinal := tabtostr(emptyword)
	fmt.Println(emptywordfinal)
	win := false
	for attemtps > 0 {
		input := getUserletter()
		if len(input) > 0 && input[0] != "" && len(input) == 1 { // test si la lettre est dans le mot
			for i := 0; i < len(wordtab); i++ {
				if string(input[0]) == wordtab[i] {
					k := i * 2
					emptyword[k] = string(wordtab[i])
					find = 1
				}
			}
			if find == 0 {
				PrintHangman("ouibonsoir", attemtps)
				attemtps--
				if attemtps == 0 {
					break
				}
				fmt.Printf("You still have %v attemtps left\n", attemtps)
			}
			find = 0
			emptywordfinal := tabtostr(emptyword)
			fmt.Println(emptywordfinal)
			finalword := ""
			for k := 0; k < len(emptyword); k++ {
				finalword += string(emptyword[k])
				k++
			}
			bateau := `			 |    |    |                 
			)_)  )_)  )_)              
		       )___))___))___)\            
		      )____)____)_____)\\
		    _____|____|____|____\\\__
           ---------\                   /---------
	           ^^^^^ ^^^^^^^^^^^^^^^^^^^^^
	          ^^^^      ^^^^     ^^^    ^^
			^^^^      ^^^`
			if finalword == word {
				if word == "bateau" {
					fmt.Println(bateau)
				}
				fmt.Println("Congrat, you win !")
				win = true
				break
			}
		} else {
			fmt.Println("Enter a valid letter")
		}
	}
	if !win {
		fmt.Printf("\nYou loose hooo ...\nthe word was %s\n", word)
	}
}
