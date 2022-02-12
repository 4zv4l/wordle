package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func rules() {
	fmt.Println(`
	Rules:
		- You have 5 guesses
		- You can only guess using 5 letters words
		- You can only guess using lowercase letters
		- Green means correct at the right position
		- Blue means correct at the wrong position
		- Red means incorrect

	Give a filename in argument to have more words
	
	Good luck!
	`)
}

// take a random word from a list of words
func getWord() string {
	var (
		words = []string{"apple", "banana", "orange", "pear", "pineapple"}
		word  string
		i     int
	)
	if len(os.Args) > 1 { // if a filename is given
		// read the file
		words, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		// split the file into words
		buff := strings.Split(string(words), "\n")
		// get a random word
		// seed the random number generator
		rand.Seed(time.Now().UnixNano())
		i = rand.Intn(len(buff))
		word = buff[i]
		// check if the word is 5 letters long
		if len(word) != 5 {
			fmt.Println("Error: the file must contain 5 letters words")
			os.Exit(1)
		}
	} else { // if no filename is given
		// take a random word from the list
		i = rand.Intn(len(words))
		word = words[i]
	}
	return word
}

func game() {
	var (
		word  string = getWord()
		guess string
		count int  = 0
		end   bool = true
	)
	for end {
		fmt.Print("> ")
		// read only five letters
		fmt.Scanf("%s", &guess)
		// flush input buffer
		// check if the word is 5 letters long
		if len(guess) != 5 {
			fmt.Println("Please enter 5 letters")
			continue
		}
		// check if the guess is correct showing colors per letter
		for i := 0; i < len(word); i++ {
			if word[i] == guess[i] { // correct at the right position
				fmt.Print("\033[32m", string(guess[i]), "\033[0m")
			} else if strings.ContainsAny(word, string(guess[i])) { // correct at the wrong position
				fmt.Print("\033[34m", string(guess[i]), "\033[0m")
			} else { // incorrect
				fmt.Print("\033[31m", string(guess[i]), "\033[0m")
			}
		}
		fmt.Println()
		// check if the whole word is guessed
		if word == guess {
			fmt.Println("\033[32m" + "You won!" + "\033[0m")
			end = false
		}
		// check if the user has 5 guesses
		if count == 5 {
			fmt.Println("\033[31m" + "You lost!" + "\033[0m")
			fmt.Println("The word was:", word)
			end = false
		}
		count++
	}
}

func main() {
	// show the rules
	rules()
	// start the game
	game()
	// exit
	fmt.Println("Thanks for playing!")
}
