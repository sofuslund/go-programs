package main

import (
    "fmt"
    "time"
    "unsafe"
    "golang.org/x/sys/windows"
)

type keyboardInput struct {
    wVk         uint16
    wScan       uint16
    dwFlags     uint32
    time        uint32
    dwExtraInfo uint64
}

type input struct {
    inputType uint32
    ki        keyboardInput
    padding   uint64
}

const (

    VK_F5 = 0x74
    VK_BACK = 0x08
    VK_CONTROL = 0x11
    VK_OEM_3 = 0xC0
    VK_OEM_6 = 0xDD
    VK_OEM_7 = 0xDE
)

var (
    user32                  = windows.NewLazyDLL("user32.dll")
    procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
    procSendInput           = user32.NewProc("SendInput")

    vKeys                   = []int{0x20, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x39, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4E, 0x4F, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, VK_OEM_3, VK_OEM_7, VK_OEM_6}
    ipt                     = input{inputType: 1, ki: keyboardInput{}}

    morseFromChar           = map[rune]string{
        'A': ".-/", 'B': "-.../", 'C': "-.-./", 'D': "-../", 'E': "./", 'F': "..-./", 'G': "--./", 'H': "..../", 'I': "../", 'J': ".---/", 'K': "-.-/", 'L': ".-../", 'M': "--/", 'N': "-./", 'O': "---/",
        'P': ".--./", 'Q': "--.-/", 'R': ".-./", 'S': ".../", 'T': "-/", 'U': "..-/", 'V': "...-/", 'W': ".--/", 'X': "-..-/", 'Y': "-.--/" , 'Z': "--../", rune(VK_OEM_3): ".-.-/", rune(VK_OEM_7): "---./", rune(VK_OEM_6): ".--.-/", '0': "-----/", '1': ".----/", '2': "..---/",
        '3': "...--/", '4': "....-/", '5': "...../", '6': "-..../", '7': "--.../", '8': "---../", '9': "----./", ' ': "/",
    }
    vKeyFromMorseLetter    = map[rune]int{
        '.': 0xBE, '-': 0xBD, '/': 0x6F,
    }
)

// Returns if a given key has been pressed since last call to this function
func isPressed(virtualKeyCode int) bool {
    r1, _, _ := procGetAsyncKeyState.Call(uintptr(virtualKeyCode));
    // Check if the least significant bit is set in the return
    // That means the key was pressed since last call to GetAsyncKeyState
    return r1 & 1 == 1
}

// Returns if a given key is down right now
func isDown(virtualKeyCode int) bool {
    isPressed(virtualKeyCode)
    return isPressed(virtualKeyCode)
}

func pressKey(vKey int) {
    ipt.ki.wVk = uint16(vKey)
    procSendInput.Call(uintptr(1), uintptr(unsafe.Pointer(&ipt)), uintptr(unsafe.Sizeof(ipt)))
    // Clear the input so that it does not get caught by the program itself and end up in an infinite loop
    isPressed(vKey)
}

func writeMorse(vKey int) {
    pressKey(VK_BACK) // Delete the caracter already typed by the user
    for _, let := range morseFromChar[rune(vKey)] {
        pressKey(vKeyFromMorseLetter[let])
    }
}

func main() {
    fmt.Println("[+] Morse Typer started ...")
    fmt.Println("[ ] type some text in any window")
    fmt.Println("[ ] the characters will be converted to morse as soon as they are typed")
    fmt.Print("[-] press f5 to stop program ")
    // Since GetAsyncKeyState return wether a key was pressed since last time the function was called we need to call it for each character one test time
    for _, vKey := range vKeys  {
        isPressed(vKey)
    }
    // While f6 is not pressed
    for !isPressed(VK_F5) {
        // Iteratate over the wanted keys which is 0 - 10 A - Z and space
        for _, vKey := range vKeys  {
            // Check if the key was pressed and control is not being held
            if isPressed(vKey) && !isDown(VK_CONTROL) {
                writeMorse(vKey)
            }
        }
        // Add a little delay so the program not uses too much cpu power
        time.Sleep(time.Millisecond * 5)
    }
}
