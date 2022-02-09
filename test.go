package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

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
			cur = 0
			return scanner.Text()
		}
		cur++
	}
	return ""
}

func main() {
	fmt.Println(readFile())
}
