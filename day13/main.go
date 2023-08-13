package main

import (
	"fmt"
)


func main() {

    test := "tuna"


    switch(test) {
        case "tuna", "ham": 
            fmt.Println("test was `",test,"`"); 
            break;

        default: fmt.Println("test was something else"); 
    }

    people := 21;
    switch {

        case people == 1: fmt.Println("Single");

        case people == 2: fmt.Println("Pair");

        case people <= 0: fmt.Println("None")

        case people > 20: fmt.Println("large Group")

        default: fmt.Println("Group")

    }

}
