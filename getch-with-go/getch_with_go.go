// +build windows,amd64

package main

import (
    "fmt"
    "golang.org/x/sys/windows"
)

func getch() rune {
    // Get a handle to the console
    var handle windows.Handle
    handle, _ = windows.GetStdHandle(windows.STD_INPUT_HANDLE)
    // Get the console mode
    var mode uint32
    windows.GetConsoleMode(handle, &mode)
    // Change the console mode so that ReadConsole does not wait for a newline
    windows.SetConsoleMode(handle, mode &^ (windows.ENABLE_LINE_INPUT | windows.ENABLE_ECHO_INPUT))
    // Get the character
    var char uint16
    var charsRead uint32
    windows.ReadConsole(handle, &char, 1, &charsRead, nil)
    // Change the console mode back to it's original state
    windows.SetConsoleMode(handle, mode)
    return rune(char)
}

func main() {
    fmt.Printf("Type some characters!")
    for {
        fmt.Printf("\nYou typed the char '%c'", getch())
    }
}
