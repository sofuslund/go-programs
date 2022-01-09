package main

import (
    "math/rand"
    "time"
    "fmt"
    "log"
)

func main() {
    log.SetPrefix("")
    log.SetFlags(0)
    // Ask user for difficulty level
    fmt.Println("1. Very easy (from 0 to 10)")
    fmt.Println("2. Easy (from 0 to 100)")
    fmt.Println("3. Hard (from 0 to 1000)")
    fmt.Println("4. Very hard (from 0 to 10000)")
    fmt.Print("Pick difficulty level: ")
    var n int
    _, err := fmt.Scan(&n)
    if err != nil {
        log.Fatal(err)
    }

    // Check if the choosen difficulty level is in bound
    if(n < 1 || n > 4) {
        log.Fatal("Difficulty level needs to be between 1 and 4")
    }
    // Set the random seed
    rand.Seed(time.Now().UnixNano())
    // Set upper bound for random number
    var ub = n * 10
    var rightNum = rand.Intn(ub)

    fmt.Println("Game start!")
    for {
        var choosenNum int
        fmt.Print("Pick a number: ")
        fmt.Scan(&choosenNum)
        if choosenNum < rightNum {
            fmt.Println("The number is higher")
        } else if choosenNum > rightNum {
            fmt.Println("The number is lower")
        } else if choosenNum == rightNum {
            fmt.Print("You won!")
            break
        }
    }
    fmt.Scanln()
    fmt.Scanln()
}
