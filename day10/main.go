package main

import (
    "fmt"
    "day10/gt"
)

var cnt1, cnt2 = 2, 5; 

func countDown(num *int, label string) gt.TickHandler {

    return func(_ float32) bool {
        fmt.Println("countdown ", label, " => ", *num);
        *num--;
        return *num > 0 
    }

}


func main() {
  
    gt.RegisterHandler(countDown(&cnt1, "cnt 1\t"));
    gt.RegisterHandler(countDown(&cnt2, "cnt 2\t"));

    l := 0;

    for  {
        l = gt.HandleTick();     

        fmt.Println("number of handlers: \t", l) 

        if l == 0  { break }

    }
    


}
