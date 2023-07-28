package main;

import ( "fmt" )

func main() {

    var x byte = 128;
    var y byte = 127;
    var z byte = 0;

    fmt.Println("xyz: ", x, y, z);
    // result => xyz:  128, 127, 0

    // Point ptrout to z  => make (*ptrout an alias for z)
    var ptrout *byte = &z;

    // Point ptrin to x  => make (*ptrin an alias for x)
    var ptrin *byte = &x;

    fmt.Println("ptrin: ", ptrin, "=", *ptrin);
    fmt.Println("ptrout: ", ptrout, "=", *ptrout);
    fmt.Println("z: ", &z, "=", z);
    // result => ptrin: [memorylocation] = 128 
    //           ptrout: [memorylocation] = 0 
    //           z: [memorylocation] = 0 


    // The following opperation would be the same as   z += x, 
    // since both pointers use an asterisk
    *ptrout += *ptrin;

    // therefore z has changed to.
    fmt.Println("ptrout: ", ptrout, "=", *ptrout);
    fmt.Println("z: ", &z, "=", z);
    // result => ptrout: [memorylocation] = 128 
    //           z: [memorylocation] = 128 


    // now point ptrin to y  => make (*ptrin an alias for y)
    ptrin = &y;

    // and run the exact same opperation as in line 30
    *ptrout += *ptrin;
    
    // therefore z has changed again.
    fmt.Println("ptrout: ", ptrout, "=", *ptrout);
    fmt.Println("z: ", &z, "=", z);
    // result => ptrout: [memorylocation] = 255 
    //           z: [memorylocation] = 255 

    fmt.Println("xyz: ", x, y, z);
    // result => xyz:  128, 127, 255 
    // Notice how z has changed, despite it never being accessed
    // directly.
    
}
