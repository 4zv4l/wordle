package main

import (
	"fmt"
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
	"log"
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

// return all the 5 length words from file
func getFileWords(f *os.File) []string {
	var scanner = bufio.NewScanner(f)
	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 {
			words = append(words, word)
		}
	}
	return words
}

// return a random 5 letters word from file
func randomFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	words := getFileWords(f)
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(words))
	return words[i]
}

// return a random word from the words list
func randomWord() string {
		// random word from list
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(words))
		return words[i]
}

// get the word to guess during the game
func getWord() string {
	var word string
	if len(os.Args) > 1 {
		word = randomFile(os.Args[1])
	} else {
		word = randomWord()
	}
	return word
}

// compare letters between user input
// and the guess word
func checkLetters(secret, guess string) {
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			fmt.Print("\033[32m", string(guess[i]), "\033[0m")
		} else if strings.ContainsAny(secret, string(guess[i])) {
			fmt.Print("\033[34m", string(guess[i]), "\033[0m")
		} else { // incorrect
			fmt.Print("\033[31m", string(guess[i]), "\033[0m")
		}
	}
	fmt.Println()
}

// get user input for the guess
func getInput() string {
	var scan = bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scan.Scan()
	if len(scan.Text()) != 5 {
		fmt.Println("Please enter 5 letters")
		return getInput()
	} else if strings.ContainsAny(scan.Text(), " ") {
		fmt.Println("No space in words please..")
		return getInput()
	}
	return scan.Text()
}

func game() {
	var (
		secret  string = getWord()
		count int  = 1
		end   bool = true
	)
	for end {
		guess := getInput()
		checkLetters(secret, guess)
		if secret == guess {
			fmt.Println("\033[32m" + "You won!" + "\033[0m")
			end = false
		}
		if count == 5 {
			fmt.Println("\033[31m" + "You lost!" + "\033[0m")
			fmt.Println("The word was:", secret)
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
