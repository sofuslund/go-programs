// +build windows,amd64

package main

import (
    "math/rand"
    "os/exec"
    "bufio"
    "time"
    "fmt"
    "os"
)

func main() {
    var reactionTime time.Duration
    var averageTime time.Duration
    var timeSum int64

    var i int64
    for i = 1; true; i++ {
        // Clear the screen
        var cmd = exec.Command("cmd", "/C", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
        // Print the data
        fmt.Print("\x1b[32m")
        fmt.Printf("█ Be ready to press enter    █ Last reaction time: %v    █ Average reaction time: %v █\n", reactionTime, averageTime)
        // Wait for random time
        var t, _ = time.ParseDuration(fmt.Sprintf("%vms", rand.Intn(3000) + 1500))
        time.Sleep(t)
        // Print to screen and start timing
        fmt.Print("\x1b[31m")
        fmt.Printf("----> boom ---->")
        var start = time.Now()
        // Wait for input
        var r = bufio.NewReader(os.Stdin)
        // Flush stdin buffer first
        fmt.Println(r.ReadString('\n'))
        // Stop timing
        reactionTime = time.Since(start)
        // Calculate average
        timeSum += reactionTime.Milliseconds()
        averageTime,_ = time.ParseDuration(fmt.Sprintf("%vms", timeSum/i))
    }
}
