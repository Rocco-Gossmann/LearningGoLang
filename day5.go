package main

import "fmt"

func main() {

	// For the counter starting at 3  count as long as the counter is less than 10  and increment the counter by 2 each cycle
	for counter := 0; counter < 10; counter++ {
		fmt.Println("Counter => ", counter)
	}

	words := [5]string{2: "hello"}

	for index := 0; index < len(words); index++ {
		fmt.Println("word at index ", index, " = ", words[index])
	}
}
