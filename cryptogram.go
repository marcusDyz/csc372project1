package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

/*
 * A struct that group all cipher keys to encrypt or decrypt the text.
 */
type cipherKeySet struct {
	encrypted_key    map[string]string
	decrypted_key    map[string]string
	user_guessed_key map[string]string
}

/*
 * A struct includes alphabet string array and a shuffled version alphbet string array
 * for the use of constructing cipherKeySet
 */
type alphabetSet struct {
	alphabet          []string
	shuffled_alphabet []string
}

// Useful information for shuffling: https://golang.cafe/blog/how-to-shuffle-a-slice-in-go.html
// initial an alphabetSet struct and shuffle the shuffled_alphabet
// and return the address of struct
/*
 * Constructor function for alphabetSet, the purpose of this function is to show
 * how to construct a struct and return its pointer via a function. The functionality of
 * this function is similar to a constructor of class in Java
 */
func constructAlphabetSet() *alphabetSet {
	alphabet, shuffled_alphabet :=
		strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", ""), strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	rand.Seed(time.Now().UnixNano())             // random generator
	rand.Shuffle(len(alphabet), func(i, j int) { // built-in Shuffle function for shuffling
		shuffled_alphabet[i], shuffled_alphabet[j] =
			shuffled_alphabet[j], shuffled_alphabet[i] // Swap the value of index i and index j of shuffled_alphabet to shuffle the alphabet
	})
	var alphabet_set = alphabetSet{alphabet, shuffled_alphabet}
	return &alphabet_set
}

/*
 * Method for cipherKeySet struct. This method initial the encrypted_key and decrypted_key fields by
 * by the passing parameter 'set', and then return both values. The purpose of this method is to
 * the method feature in Go language which is similar to a private setter method for a class in Java.
 */
func (cipher cipherKeySet) initKeys(set *alphabetSet) (map[string]string, map[string]string) {
	alphabet := set.alphabet
	shuffled_alphabet := set.shuffled_alphabet
	encrypted_key := make(map[string]string)
	decrypted_key := make(map[string]string)
	for i := 0; i < 26; i++ {
		encrypted_key[alphabet[i]] = shuffled_alphabet[i]
		decrypted_key[shuffled_alphabet[i]] = alphabet[i]
	}
	return encrypted_key, decrypted_key
}

// use the alphabet and the shuffled_alphabet to initial a cipherKeySet struct
// and return the struct
/*
 * The purpose of this function is Similar to constructAlphabetSet. In addition,
 * this construct function uses method feature for struct.
 */
func constructKeySet(set *alphabetSet) *cipherKeySet {
	encrypted_key, decrypted_key, guessed_key := make(map[string]string), make(map[string]string), make(map[string]string)
	var cipher_key_set = cipherKeySet{encrypted_key, decrypted_key,
		guessed_key}
	cipher_key_set.encrypted_key, cipher_key_set.decrypted_key = cipher_key_set.initKeys(set) // uses Method feature

	return &cipher_key_set
}

// use the cipherKeySet to encrypt the text
// return the encrypted text string
// line 67 uses the technique mentioned by the top answer in https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
func encryptText(cipher cipherKeySet, original_text string) string {
	result := ""
	encrypted_key := cipher.encrypted_key
	// no original_text string intialized yet.
	for i := 0; i < len(original_text); i++ {
		if val, ok := encrypted_key[string(original_text[i])]; ok {
			result += val
		} else {
			result += string(original_text[i])
		}
	}
	return result
}

// decrypt the encrypted_text based on user's input
// return the user decrypted text string
func userGuessedText(cipher cipherKeySet, encrypted_text string) string {
	guessed_key := cipher.user_guessed_key
	result := ""
	for i := 0; i < len(encrypted_text); i++ {
		if val, ok := guessed_key[string(encrypted_text[i])]; ok {
			result += val
		} else {
			if !unicode.IsLetter(rune(encrypted_text[i])) {
				result += string(encrypted_text[i])
			} else {
				result += " "
			}
		}
	}
	return result
}

// '_' in Golang represents a "blank identifier", it avoids having to declare all the variables for the returns values.
// In other word, it discards the no uss return value
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
	original_text := strings.ToUpper(readFile())
	alphabet := constructAlphabetSet()
	cipher := constructKeySet(alphabet)
	encrypted_text := encryptText(*cipher, original_text)
	fmt.Println(encrypted_text)
	fmt.Println(userGuessedText(*cipher, encrypted_text))
	// Keep asking user input until "exit" has been typed
	for {
		consoleReader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter the letter to replace: ")
		replace_letter, _ := consoleReader.ReadString('\n')
		replace_letter = strings.ToUpper(replace_letter)
		replace_letter = strings.TrimSuffix(replace_letter, "\n")

		// check if the input is exit
		if strings.HasPrefix(replace_letter, "EXIT") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		fmt.Print("Enter its replacement: ")
		guess_letter, _ := consoleReader.ReadString('\n')
		guess_letter = strings.ToUpper(guess_letter)
		guess_letter = strings.TrimSuffix(guess_letter, "\n")

		if strings.HasPrefix(replace_letter, "EXIT") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}
		cipher.user_guessed_key[replace_letter] = guess_letter
		fmt.Println(encrypted_text)
		user_attempt := userGuessedText(*cipher, encrypted_text)
		fmt.Println(user_attempt)
		// compare the user attempt is equal to the original text
		if user_attempt == original_text {
			fmt.Print("You Win!\n")
			os.Exit(0)
		}
	}
}
