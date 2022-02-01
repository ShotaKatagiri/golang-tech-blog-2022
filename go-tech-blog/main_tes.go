package main

import "fmt"


func main() {
    fmt.Print("Hello, ")             // 改行されない
    fmt.Println("世界")               // 改行される
    fmt.Println("A", 100, true, 1.5) // スペース区切りで表示される
}