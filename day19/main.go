package main

import "fmt"

func youMayPanicNow(num uint) uint {
    if num == 5 {
        panic("PAAAAAAAAANIC!!!!!!! *runs in circles*");
    }
    return num;
}

func recoverFromPanic() {
    if hadPanic := recover(); hadPanic != nil {
        fmt.Printf("Recover from Panic: %v\n", hadPanic);    
        // What value created the panic can not be defined at this point 
        // (Panic would need to return an interface or struct and the 
        //  function that caused the panic would need an extra param too. 
        //  I'm just to lazy to do that for this example)
    }
}



// We use the "Comma, OK idium" to define if a panic happened 
// (The last param will be true if no panic happened)
func giveMeANumber_5WillPanic(checkA, checkB uint) (int, int, bool) {

    defer recoverFromPanic(); // It is important, that a 
                              // panic recovery function is called via defer

    // Defining some vars for the return (Since this function has no named returns)
    //------------------------------------------------------
    a,b,ok := -1,-1,true; 
            // /\-- Default Values serve no purpouse but to demonstrate that
            //      in a function without named returns, that experienced a panic 
            //      it does not matter what the return values are
            //   
            //      only if NO panic happened will these 3 values be returned
            //      So we can set ok to true here, but the function will only 
            //      return ok if no panic happens 
            //
            //      otherwise it returns false because
            //      that is the default value for undefined bool types.

    a = int(youMayPanicNow(checkA))
    b = int(youMayPanicNow(checkB))

    return a, b, ok // <-- The return step is skipped during a panic
                    //     So all paremters defined in the function return-
                    //     signature are set to their default value
}

func giveMeANumber_5WillPanic_RecoveryViaNamedReturns(checkA, checkB uint) (a int, b int, ok bool) {

    a = -1; // <-- define default values that 
    b = -1; //     are used in case of a panic

    ok = true; // <-- set to false in case of panic

    // You can also define panic handlers as inline functions in that case
    // The handler gets access to all values in its defining scope
    defer func() { // <- In case any function after this definition panics
        if hadPanic := recover(); hadPanic != nil {
            fmt.Printf("Recover from Panic: %v\n", hadPanic);    
            ok = false;
        }
    } ()

    a = int(youMayPanicNow(checkA))
    b = int(youMayPanicNow(checkB))

    return //<- named returns are filled in during function execution
           //   So unlike `giveMeANumber_5WillPanic` here, we get the default
           //   values, we defined earlier since the return vars are not filled
           //   in by the return statement (which was skipped by the panic)
}

func main() {


    fmt.Println("giveMeANumber_5WillPanic(6, 10)\n================================================================================")
    a,b,ok := giveMeANumber_5WillPanic(6, 10)
    if ok != true {
        fmt.Println("\n\nDiscovered Panic in 'giveMeANumber_5WillPanic'", ok);
    }
    fmt.Printf("A,B running without panic is: %v, %v\n\n\n", a,b);



    fmt.Println("giveMeANumber_5WillPanic(10, 5)\n================================================================================")
    a,b,ok = giveMeANumber_5WillPanic(10, 5)
    if ok != true {
        fmt.Println("Discovered Panic in 'giveMeANumber_5WillPanic'");
    }
    fmt.Printf("A,B is: %v, %v\n\n\n", a,b);



    fmt.Println("giveMeANumber_5WillPanic_RecoveryViaNamedReturns(5, 10)\n================================================================================")
    a,b,ok = giveMeANumber_5WillPanic_RecoveryViaNamedReturns(5, 10)
    if ok != true {
        fmt.Println("Discovered Panic in 'giveMeANumber_5WillPanic_RecoveryViaNamedReturns'");
    }
    fmt.Printf("A,B is: %v, %v\n\n\n", a,b);

    fmt.Println("giveMeANumber_5WillPanic_RecoveryViaNamedReturns(9, 5)\n================================================================================")
    a,b,ok = giveMeANumber_5WillPanic_RecoveryViaNamedReturns(9, 5)
    if ok != true {
        fmt.Println("Discovered Panic in 'giveMeANumber_5WillPanic_RecoveryViaNamedReturns'");
    }
    fmt.Printf("A,B is: %v, %v\n", a,b);

}
