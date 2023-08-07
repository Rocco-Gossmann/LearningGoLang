package main

import "fmt"

var (
	a byte = 1;
	b byte = 2;

	c *byte = new(byte);
)

func main() {
	fmt.Printf("a: at address %p of type %T \thas value %#x\n", &a, a, a)
	fmt.Printf("b: at address %p of type %T \thas value %#x\n", &b, b, b)
	fmt.Printf("c: at address %p of type %T \thas value %#x\n", &c, c, c)
}

