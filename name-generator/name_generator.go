package main

import (
    "fmt"
    "log"
    "strings"
    "math/rand"
)

func main() {
    const (
        // There are multiple of some letters since they are more common
        VOWELS = "aaeeeiouy"
        CONSONANTS = "bcdfghjklmnnpqrrrsttvwxz"
    )

    // Get name length
    var length int
    fmt.Print("Enter name length: ")
    _,err := fmt.Scanf("%d\n", &length)
    if err != nil {
        log.Fatal(err)
    }
    // Get amount of names to generate
    var amount int
    fmt.Print("Enter amount of names to be generated: ")
    _,err = fmt.Scanf("%d\n", &amount)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Generating names...")

    // Generate names
    for i := 0; i < amount; i++ {
        var name = ""
        var startWithVowel = rand.Intn(1)
        for j := 0; j < length; j++ {
            if (j+startWithVowel) % 2 == 0{
                name += string(CONSONANTS[rand.Intn(len(CONSONANTS))])
                // There should be more cononants in the names
                if rand.Intn(3) == 0 {
                    j--
                }
            } else {
                name += string(VOWELS[rand.Intn(len(VOWELS))])
            }
        }
        // Print the name capitalized
        fmt.Println(strings.Title(name))
    }
}
