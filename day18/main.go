package main

type Color uint8
var (
    Red   Color = 1
    Green Color = 2
    Blue  Color = 3
)

var NotAColor uint8 = 128;

func giveMeAColor(c *Color) { }  

func main() {

    giveMeAColor(&Red);
    giveMeAColor(&NotAColor);
    giveMeAColor(&Color(128));

}
