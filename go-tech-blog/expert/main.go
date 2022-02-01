package main

import (
	"fmt"
	"time"
)

type Person struct {
    Name string
    Age int
}



func main() {

    var d *Person
    d = &Person {
        Name: "shotaa",
        Age: 27,
    }
    tb := time.Now()
    fmt.Printf("d :%+v\n", d)             // 改行されない
    fmt.Println(tb)                // 改行される
    fmt.Println("A", 100, true, 1.5) // スペース区切りで表示される
}