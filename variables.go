package main

import "fmt"

//This syntax can be used both inside functions and at the package level
var z string = "blablabla"

func main() {

	// this syntax can only be used inside functions, not at the package level
    var a = "initial"
    fmt.Println(a) //initial

    var b, c int = 1, 2
    fmt.Println(b, c) //1 2

    var d = true
    fmt.Println(d) //true

    var e int
    fmt.Println(e) //0

    f := "apple"
    fmt.Println(f) //apple

	fmt.Println(z) //blablabla
}