package main

import (
    "strings"
    "code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    var words []string = strings.Split(s," ");
   
    for i := range words {
        m[words[i]] = m[words[i]] + 1
    }
    return m
}

func main() {
    wc.Test(WordCount)
}


