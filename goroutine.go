package main

import (
    "fmt"
    "time"
)

func concatStrings(s []string, channel chan string) {
    concat := ""
    for i := 0; i < len(s); i++ {
        concat += s[i]
    }
    channel <- concat
}

func printNums(n int, m int) {
    for i := n; i <= m; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(i)
    }
}

func main() {
    go printNums(1, 5)
    printNums(6, 10)

    s := []string{"testing ", "go ", "channels"}
    channel := make(chan string)

    go concatStrings(s, channel)

    finalString := <-channel
    fmt.Println("Using channel to direct goroutine on concatStrings")
    fmt.Println(finalString)
}
