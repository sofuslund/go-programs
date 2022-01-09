package main

import (
    "os"
    "fmt"
    "log"
    "regexp"
)

func main() {
    log.SetPrefix("")
    log.SetFlags(0)
    if len(os.Args) < 2 {
        log.Fatal("Usage: password_check PASSWORD\n\tPASSWORD: A string surrounded by quotes if spaces is in")
    }
    var password = os.Args[1]
    fmt.Printf("Test no. | Test                                             | Passed\n")
    fmt.Printf("-------------------------------------------------------------------\n")
    // If it's length is at least 8
    var matched, _ = regexp.MatchString(`^.{8,}$`, password)
    fmt.Printf("1        | Contains at least 8 characters                   | %v\n", matched)
    // If it contains at least 1 lowercase letter
    matched, _ = regexp.MatchString(`^(.*[a-z].*[A-Z].*|.*[A-Z].*[a-z].*)$`, password)
    fmt.Printf("2        | Contains both uppercase and lowercase characters | %v\n", matched)
    // If it contains at least 1 digit
    matched, _ = regexp.MatchString(`^.*\d.*$`, password)
    fmt.Printf("3        | Contains digits                                  | %v\n", matched)
}
