// Program that can run different sorting algorithms and speed test them
// The package sortingalgorithms implements the algorithms while this main package parses the arguments and times the algorithms
// You run the program by typing:
// run-sorting-alg.exe ALGORITHMS on windows in command prompt or:
// ./run-sorting-alg ALGORITHMS on linux or mac in terminal
// where ALGORITHMS is sorting algorithms seperated by spaces
// Each algorithm can be one of the following:
// selection_sort
// bubble_sort
// and can be abbreviated to their first letter
// For help use the --help or -h option
package main

import (
    "os"
    "fmt"
    "log"
    "time"
    "math/rand"
    algs "example.com/run-sorting-alg/sortingalgorithms"
)

const usageMsg =
`Usage: run-sorting-alg [--help] ALGORITHMS

optional arguments:
    -h, --help: Shows this message and quits
positional arguments:
    ALGORITHMS: Space seperated algorithms. An algorithm can be one of the following:
    'selection_sort', 'bubble_sort', 'insertion_sort', 'merge_sort', 'quick_sort', 'all'

The algorithm names can also be abbreviated to their first letter
`

func main() {
    // ==================================================
    // INITIALIZE                                       |
    // ==================================================

    // Set the logging options as the prefix and flags
    log.SetPrefix("")
    log.SetFlags(0)

    rand.Seed(time.Now().Unix())
    var arr = rand.Perm(400000)

    // ==================================================
    // PARSE ARGUMENTS                                   |
    // ==================================================

    // Print usage message and quit if no algorithm was specified
    if len(os.Args) < 2 {
        log.Fatal(usageMsg)
    }

    // Operate on each argument one by one
    for _, arg := range os.Args[1:] {
        // Check if the argument is optional by testing if it has a leading dash
        if arg[0] == '-' {
            // Check if the option is long by checking if it has one more dash
            if arg[1] == '-' {
                switch arg {
                case "--help":
                    log.Fatal(usageMsg)
                default:
                    log.Fatal("unknown option: " + arg)
                }
            } else { // Else the option has only one dash and is short
                for _, shrtOpt := range arg[1:] {
                    switch shrtOpt {
                    case 'h':
                        log.Fatal(usageMsg)
                    default:
                        log.Fatal("unknown option: -" + string(shrtOpt))
                    }
                }
            }
            continue
        }
        // If the argument is a positional argument
        fmt.Print("BEFORE: ", arr, "\n")
        fmt.Print("AFTER: ")
        before := time.Now()
        switch arg {
            case "s", "selection_sort":
                fmt.Println(algs.SelectionSort(arr))
            case "b", "bubble_sort":
                fmt.Println(algs.BubbleSort(arr))
            case "i", "insertion_sort":
                fmt.Println(algs.InsertionSort(arr))
            case "m", "merge_sort":
                fmt.Println(algs.MergeSort(arr))
            case "q", "quick_sort":
                fmt.Println(algs.QuickSort(arr))
            case "a", "all":
                fmt.Println(algs.SelectionSort(arr))
                fmt.Println(algs.BubbleSort(arr))
                fmt.Println(algs.InsertionSort(arr))
                fmt.Println(algs.MergeSort(arr))
                fmt.Println(algs.QuickSort(arr))
            default:
                log.Fatal("unrecognized algorithm: " + arg)
        }
        fmt.Println(time.Since(before))
    }
}
