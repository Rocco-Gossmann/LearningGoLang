package main

import "fmt"
import "strings"

func transformString(input string) (string, string){
    if len(input) >= 3 {
        return strings.ToUpper(input), ""
    } else {
        return input                 , "'"+input+"' is less than 3 chars"
    }
}

func printResult(result string, err string) {
    if err == "" {
        fmt.Println("transformed string: ", result);
    } else {
        fmt.Println("failed to transform String: ", err);
    }
} 

func main() {
    
    printResult(transformString("hi")) 

    printResult(transformString("hello")) 

}
