package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type cipherKeySet struct {
	encrypted_key    map[string]string
	decrypted_key    map[string]string
	user_guessed_key map[string]string
}

type alphabetSet struct {
	alphabet          []string
	shuffled_alphabet []string
}

// useful information for shuffling: https://golang.cafe/blog/how-to-shuffle-a-slice-in-go.html
// initial an alphabetSet struct and shuffle the shuffled_alphabet
// and return the struct
func constructAlphabetSet() *alphabetSet {

}

// use the alphabet and the shuffled_alphabet to initial a cipherKeySet struct
// and return the struct
func constructKeySet(set alphabetSet) *cipherKeySet {

}

// use the cipherKeySet to encrypt the text
// return the encrypted text string
func encryptText(cipher cipherKeySet) string {

}

// decrpyt the encrypted_text based on user's input
// return the user decrypted text string
func decryptText(cipher cipherKeySet) string {

}

func readFile() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randIndex := r1.Intn(5)

	dir, _ := os.Getwd()

	file, err := os.Open(dir + "/quote.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cur := 0
	for scanner.Scan() {
		if cur == randIndex {
			return scanner.Text()
		}
		cur++
	}
	return ""
}

func main() {
	fmt.Println(strings.ToUpper(readFile()))

}
