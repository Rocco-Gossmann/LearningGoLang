package main

import "fmt"

func printInfo(label string, baseArr *[5]byte, slice *[]byte) {

    fmt.Print("\n", label, "\n==============================\n")

    fmt.Printf("Arr:\tlen %d cap %d\n", len(baseArr), cap(baseArr))
    fmt.Printf("slice:\tlen %d cap %d\n", len(*slice), cap(*slice))
    fmt.Print("\n")
    fmt.Print("Arr:\t");

    for _, b := range baseArr { fmt.Printf("%d\t\t", b); }
    fmt.Print("\nArr:\t")
    for i := range baseArr { fmt.Printf("%p\t", &baseArr[i]); }

    fmt.Print("\n\n")

    fmt.Print("Slice:\t");
    for _, b := range *slice { fmt.Printf("%d\t\t", b); }
    fmt.Print("\nSlice:\t")
    for i := range *slice { fmt.Printf("%p\t", &(*slice)[i]); }

    fmt.Println("\n------------------------------")
}

func main() {
  
    var baseArr = [5]byte{1, 2, 4, 8, 16}
    var slice = baseArr[1:3];

    printInfo("initial state", &baseArr, &slice);

    slice = append(slice, 32)
    printInfo("after append, without breaking the capacity", &baseArr, &slice);

    slice = append(slice, 64, 128)
    printInfo("after breaking the capacity", &baseArr, &slice);

}
