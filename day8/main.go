package main

import (
	"day8/speak"
	"day8/speak/words"
	"fmt"
    "math/rand"
)

func fnc1() { fmt.Println("func 1"); }
func fnc2() { fmt.Println("func 2"); }

type handlerFunction func()
type bar string;

func main() {

    // Function pointer example
    fncArr := [4]handlerFunction{ fnc1 ,fnc2 ,words.Speak ,speak.SayHello }

    // Pic a random entry in fncArr
    var fncIndex =  int(rand.Float32() * float32(len(fncArr)-1) + 0.5);

    fncArr[fncIndex]();
  
    // Custom Type example
    var foo bar = "hello you";
    fmt.Printf("foo is type: %T\n", foo);

    fmt.Println("\n\nArray-Slice") 
    fmt.Println("================================================");

    var sliceBase = [5]byte{1, 2, 3, 4, 5};
    numSlice := sliceBase[:];
    numSlice[1] = 111;

    fmt.Println("var sliceBase = [5]byte{1, 2, 3, 4, 5};");
    fmt.Println("var numSlice []byte = sliceBase[1:3];");
    fmt.Println("\n------------------------------------------------");
    fmt.Printf("sliceBase\tis %T  of size %d with capacity %d => %v\n", sliceBase, len(sliceBase), cap(sliceBase), sliceBase);
    fmt.Printf("numSlice\tis %T  of size %d with capacity %d => %v\n", numSlice, len(numSlice), cap(numSlice), numSlice);
//  fmt.Println("\nsliceBase's capacity will always be its size and that can never change");
//  fmt.Println("\nnumSlice's size is 2 because we started at index 1 and ended BEFORE index 3 (that is what the [1:3] means)");
//  fmt.Println("Its capacity however is 4 because to the end of the arrays, that this slice is lookit at, there are 2 more elements");
//  fmt.Println("");
//  fmt.Println("     |  1,  [2,  3,]  4,  5  |   <-- Base Array (sliceBase)");
//  fmt.Println("             /\\     /\\      /\\ ");
//  fmt.Println("             |      |       | - Array capacity End");
//  fmt.Println("             |       -Slice |");
//  fmt.Println("             |        end   |");
//  fmt.Println("              -Slice Start  |");
//  fmt.Println("             |              |");
//  fmt.Println("             ----------------");
//  fmt.Println("                  |- Slice Capacity");




    numSlice = append(numSlice, 6);

    fmt.Println("\n------------------------------------------------");
    fmt.Println("after appending 1 element  ")
    fmt.Println("------------------------------------------------");
    fmt.Printf("sliceBase\tis %T  of size %d with capacity %d => %v\n", sliceBase, len(sliceBase), cap(sliceBase), sliceBase);
    fmt.Printf("numSlice\tis %T  of size %d with capacity %d => %v\n", numSlice, len(numSlice), cap(numSlice), numSlice);
    fmt.Println("\n(notice how the sliceBases last value changed)");





    numSlice[1] = 255 
    fmt.Println("\n------------------------------------------------");
    fmt.Println("after chaning the 2nd element of the slice : ")
    fmt.Println("------------------------------------------------");
    fmt.Printf("sliceBase\tis %T  of size %d with capacity %d => %v\n", sliceBase, len(sliceBase), cap(sliceBase), sliceBase);
    fmt.Printf("numSlice\tis %T  of size %d with capacity %d => %v\n", numSlice, len(numSlice), cap(numSlice), numSlice);

    numSlice = append(numSlice, 7, 8, 9, 10, 11, 12, 13, 14, 15);
    numSlice[0] = 127

    fmt.Println("\n------------------------------------------------");
    fmt.Println("after extending the slice:")
    fmt.Println("------------------------------------------------");
    fmt.Printf("sliceBase\tis %T  of size %d with capacity %d => %v\n", sliceBase, len(sliceBase), cap(sliceBase), sliceBase);
    fmt.Printf("numSlice\tis %T  of size %d with capacity %d => %v\n", numSlice, len(numSlice), cap(numSlice), numSlice);

    fmt.Println("------------------------------------------------\n\n");


    fmt.Println("\n\nStandalone-Slice") 
    fmt.Println("================================================");

    var soSlice = []byte{1, 2, 3, 4, 5}
    fmt.Println("var soSlice = []byte{1, 2, 3, 4, 5}");
    fmt.Println("------------------------------------------------");
    fmt.Println("after definition");
    fmt.Println("------------------------------------------------");
    fmt.Printf("soSlice\tis %T  of size %d with capacity %d => %v\n", soSlice, len(soSlice), cap(soSlice), soSlice);

    fmt.Println("\n\nSlice via make") 
    fmt.Println("================================================");

    //var makeSlice = make([]byte, 2, 5);
    var makeSlice []byte = []byte{}

    
    fmt.Println("------------------------------------------------");
    fmt.Println("after definition");
    fmt.Println("------------------------------------------------");
    fmt.Printf("makeSlice\tis %T  of size %d with capacity %d => %v\n", makeSlice, len(makeSlice), cap(makeSlice), makeSlice);
}
