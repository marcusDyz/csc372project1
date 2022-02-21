package main

import (
    "fmt"
    "time"
)

/*
concatStrings function,
concats a list of strings into one single string.
@ param: s, a list of strings.
@ return: returns the concatenated string, sends the
value into channel which receives it.
*/
func concatStrings(s []string, channel chan string) {
    concat := ""
    for i := 0; i < len(s); i++ {
        concat += s[i]
    }
    channel <- concat
}

/*
printNums function,
A simple function that prints numbers from
value n to value m. Uses time.Sleep to slow the
time between each iteration.
*/
func printNums(n int, m int) {
    for i := n; i <= m; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(i)
    }
}

func main() {
    go printNums(1, 5) // uses 'go' to start goroutine.
    printNums(6, 10)
    // Both functions will run simultaneously, as seen in
    // the output.

    s := []string{"testing ", "go ", "channels"}
    channel := make(chan string) // initiate string channel.

    go concatStrings(s, channel) // uses 'go' to start goroutine.

    finalString := <-channel // sends result to channel.
    fmt.Println("Using channel to direct goroutine on concatStrings")
    fmt.Println(finalString)
}
