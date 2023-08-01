package main

import "fmt"


func main() {

// Empty array defintion
    var wordlist [2]string

//  var      wordlist    [             2        ]    string
//  the variable wordlist is an array with 2 elements of type string

// If you want to infere the number of elements by the definitions 
    words := []string{ "hello", "world" }


    fmt.Printf("Values: %v, %v\n", words, wordlist);
}

