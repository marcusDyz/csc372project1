package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func readFile() string {
	randIndex := rand.Intn(5)
	file, err := os.Open("/Users/dengmarcus/Desktop/quote.txt")
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
	fmt.Println(readFile())
}
