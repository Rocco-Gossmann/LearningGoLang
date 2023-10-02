package main

import "fmt"

//==============================================================================
// Structs
//==============================================================================
type Coords struct {
    x,y int32
}
type Color struct {
    r, g, b uint8
}

type Pixel struct {
    x, y int32
    coordinates *Coords

    r, g, b uint8
    color *Color
}

//==============================================================================
// Functions
//==============================================================================

func usesAPixel(opts Pixel) { 
    fmt.Println("Pixel: ", opts) 

}

func usesAStruct(opts struct{
    myarg1 string
    myarg2 int
}) { fmt.Println("Struct: ", opts) }

//==============================================================================
// Main
//==============================================================================
func main() {

    var red = Color{255, 0, 0};
    var coordinates = Coords{20, 20};

    usesAPixel(Pixel{10, 10, nil, 51, 51, 51, nil});
    usesAPixel(Pixel{color: &red});
    usesAPixel(Pixel{coordinates: &coordinates});
    usesAPixel(Pixel{coordinates: &coordinates, r: 128, g: 128, b: 128});
    usesAPixel(Pixel{coordinates: &Coords{x: 20}});
    usesAPixel(Pixel{coordinates: &Coords{x: 30}});

}
