package expert

import (
    "fmt"
    "time"
)


func main() {
    tb := time.Now()
    fmt.Print("Hello, ")             // 改行されない
    fmt.Println(tb)               // 改行される
    fmt.Println("A", 100, true, 1.5) // スペース区切りで表示される
}

