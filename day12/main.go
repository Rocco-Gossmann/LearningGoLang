package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func printInfo(label string, baseArr *[5]byte, slice1 *[]byte, slice2 *[]byte) {

    fmt.Print("\n", label, "\n==============================\n")

    fmt.Printf("Arr:  %p \tlen %d cap %d\n"   , baseArr, len(baseArr), cap(baseArr))
    fmt.Print("\n")
    
    fmt.Printf("slice1: %p \tlen %d cap %d\n", slice1 , len(*slice1), cap(*slice1))
    fmt.Printf("%#x\n\n", (*reflect.SliceHeader)(unsafe.Pointer(slice1)));
    fmt.Printf("slice2: %p \tlen %d cap %d\n", slice2 , len(*slice2), cap(*slice2))
    fmt.Printf("%#x\n\n", (*reflect.SliceHeader)(unsafe.Pointer(slice2)));
    fmt.Print("\n")
    fmt.Print("Arr:\t");

    for _, b := range baseArr { fmt.Printf("%d\t\t", b); }
    fmt.Print("\nArr:\t")
    for i := range baseArr { fmt.Printf("%p\t", &baseArr[i]); }

    fmt.Print("\n\n")

    fmt.Print("Slice1:\t");
    for _, b := range *slice1 { fmt.Printf("%d\t\t", b); }
    fmt.Print("\nSlice1:\t")
    for i := range *slice1 { fmt.Printf("%p\t", &(*slice1)[i]); }

    fmt.Print("\n\n")
    fmt.Print("Slice2:\t");
    for _, b := range *slice2 { fmt.Printf("%d\t\t", b); }
    fmt.Print("\nSlice2:\t")
    for i := range *slice2 { fmt.Printf("%p\t", &(*slice2)[i]); }

    fmt.Println("\n------------------------------")
}

func main() {
  
    var baseArr = [5]byte{1, 2, 4, 8, 16}
    var slice1 = baseArr[1:3];
    var slice2 = baseArr[2:];

    printInfo("initial state", &baseArr, &slice1, &slice2);

    slice1 = append(slice1, 32)
    printInfo("after append", &baseArr, &slice1, &slice2);

    baseArr[2] = 64;
    printInfo("after changing base", &baseArr, &slice1, &slice2);

    slice1 = slice1[1:]
    printInfo("after reslice slice1", &baseArr, &slice1, &slice2);

}
