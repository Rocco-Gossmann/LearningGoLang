package main

type Color uint8
const (
    Red Color = iota 
    Green
    Blue
)

const NotAColor uint8 = 128;

func giveMeAColor(c Color) {}  

func main() {

    giveMeAColor(Red);
    giveMeAColor(NotAColor);

}
