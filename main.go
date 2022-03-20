package main

import (
	"fmt"
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/chzyer/readline"
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
func getWordsFile(f *os.File) []string {
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
func randomWordFile(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	words := getWordsFile(f)
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
		word = randomWordFile(os.Args[1])
	} else {
		word = randomWord()
	}
	return word
}

// compare letters between user input
// and the guess word
func checkLetters(secret, guess string) {
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] { // right place
			fmt.Print("\033[32m", string(guess[i]), "\033[0m")
		} else if strings.ContainsAny(secret, string(guess[i])) { // wrong place
			fmt.Print("\033[34m", string(guess[i]), "\033[0m")
		} else { // incorrect
			fmt.Print("\033[31m", string(guess[i]), "\033[0m")
		}
	}
	fmt.Println()
}

// check if input contains spaces
// or if input is more/less than 5 characters long
func checkInput(input string) bool {
	if len(input) != 5 {
		fmt.Println("Please enter 5 letters")
		return false
	} else if strings.ContainsAny(input, " ") {
		fmt.Println("No space in words please..")
		return false
	}
	return true
}

// get user input for the guess
func getInput() string {
	reader, err := readline.New("> ")
	if err != nil { return getInput() }
	input, err := reader.Readline()
	if err != nil { return getInput() }
	if !checkInput(string(input)) { return getInput() }
	return string(input)
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
		if count == 5 && end {
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
