package main

import "fmt"

type FncPointer func() FncPointer

var run = 0

func pointToMe() FncPointer {
	if run < 10 {
		run++
	fmt.Println("Running:", run)
return pointToMe
	}
	return nil
}

func useFnc(f FncPointer) {
	for f != nil {
		fmt.Printf("fncAddr: %p\n", f)
		f = f()
	}
}

func main() {
	useFnc(pointToMe)
}
