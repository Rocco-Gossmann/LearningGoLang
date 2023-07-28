package main;

import "fmt"

func main() {

    var (
        firstname string = "John"
        lastname  string = "Doe"
    )

    const (
        pi float64 = 3.14159265359
        pi2 float32 = 6.28318530718
    ) 

    var count = 0

    fmt.Printf("%s\n%s\n%f\n%f\n%d\n\n", firstname, lastname, pi, pi2, count)

}
