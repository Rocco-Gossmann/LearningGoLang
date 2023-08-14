package main

import "fmt"



type characterList map[string]uint8


func printCharacters(chars *characterList) {
    fmt.Printf("%d character(s) in list:\n--------------------\n", len(*chars)) 
    for name, age := range *chars {
        fmt.Printf("%s: \t%d year(s) old\n", name, age);
    }

    fmt.Println();
}

func exists(chars *characterList, charName string) (bool,string) {

    if age, ok := (*chars)[charName]; ok {
        fmt.Printf("%s:\t %d years(s) old\n", charName, age);
        return ok, charName;
    } else {
        fmt.Printf("%s is not part of the list (%v)\n", charName, ok);
        return ok, charName;
    }

}



func main() {

    characters := characterList{
        "Doppo": 56,
        "Katsumi": 21,
    }

    printCharacters(&characters);
    ok, _ := exists(&characters, "Katou");

    var name string;
    fmt.Printf("%t\n", ok)
    ok, name = exists(&characters, "Doppo");
    fmt.Printf("%s exists: %t\n", name, ok)


}
