// +build windows,amd64

package main

import (
    "os"
    "fmt"
    "syscall"
)

var (
	kernel32Dll    *syscall.LazyDLL  = syscall.NewLazyDLL("Kernel32.dll")
	pSetConsoleMode *syscall.LazyProc = kernel32Dll.NewProc("SetConsoleMode")
)

func enableVirtualTerminalProcessing() {
	const ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

    	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING

	ret, _, err := pSetConsoleMode.Call(uintptr(syscall.Stdout), uintptr(mode))
	if ret == 0 {
		fmt.Fprintln(os.Stderr, err)
	}
}

const(
    ESC="\x1b"
    CSI="\x1b["
    OSC="\x1b]"
    BEL="\x07"
)

func printRGB(r, g, b int) {
    fmt.Printf("%s38;2;%d;%d;%dm", CSI, r, g, b)
    fmt.Print("█")
}

func printHUE(step_size int) {
    if(255 % step_size != 0) {
        fmt.Fprintln(os.Stderr, "step size should be a factor of 255")
    }
    var rgb = [3]int{255,0,0}
    var rise = [3]int{1, 2, 0}
    var fall = [3]int{0, 1, 2}

    for i := 0; i < 3; i++ {
        // Make a pointer to the color channel that is currently getting modified for easier syntax
        var c = &rgb[rise[i]]
        // Get the color channel that should rise and rise it gradually
        for *c = 0; *c < 255; *c += step_size {
            // Print a block character in the colors
            printRGB(rgb[0], rgb[1], rgb[2])
        }
        fmt.Println("")
        // Get the color channel that should fall and fall it gradually
        c = &rgb[fall[i]]
        for *c = 255; *c > 0; *c -= step_size {
            // Print a block character in the colors
            printRGB(rgb[0], rgb[1], rgb[2])
        }
	if i != 2 {
	    fmt.Println("")
        }
    }
}

func main() {
    // We need to enable console virtual terminal processing to use VT100 escape codes on windows
    enableVirtualTerminalProcessing()
    // Change the console title to "Colors!"
    fmt.Printf("%s2;Colors!%s", OSC, BEL)

    var factorsOf255 = [8]int{1, 3, 5, 15, 17, 51, 85, 255}
    printHUE(factorsOf255[1])

    // Reset the console colors
    fmt.Print(CSI+"m")
    fmt.Scanln()
}
