package main

import (
    "fmt"
    "math/rand"
    "time"
    "strings"
)

func main() {

    myTime := time.Now().Unix();
    fmt.Println("my Timestamp: ", myTime)
    rand.Seed(myTime)

	words := [1]string{"hello"}
    word := "world"
    fmt.Printf("%p: %v\n%p: %v\n", &words[0], words, &word, word);

    rnd := int(rand.Float32() * 1024)
    fmt.Println(rnd);
    
    words[0] = strings.Repeat("a", rnd)

    fmt.Printf("%p: %v\n%p: %v\n", &words[0], words, &word, word);

}



