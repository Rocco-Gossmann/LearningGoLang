package main

import (
	"fmt"
)

func main() {


    stringList := []string{"santa", "clause", "santa", "dorian"};
    fmt.Printf("StringList: %v\n", stringList);


    filterMap := map[string]bool{};
    for _, word := range stringList {
        filterMap[word] = true
    }

    stringList = stringList[0:0];
    for word := range filterMap {
        stringList = append(stringList, word);
    }

    fmt.Printf("StringList: %v\n", stringList);


    characterAges := map[string]uint8 {
        "Doppo": 56,
        "Katsumi": 21,
    }


    char, ok := characterAges["Katou"];

    fmt.Printf("%d, %v\n", char, ok);

    

}
