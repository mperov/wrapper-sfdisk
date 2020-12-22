package main

import (
    "fmt"
    "os"
)

const nameUtility = "wrapper-disk"

func greet() {
    fmt.Println(fmt.Sprintf("\nWelcome to %v!", nameUtility))
    fmt.Println("Changes which is written on disk can destroy dates on your disk!")
    fmt.Println("Please, be careful using this utility!")
}

func usage() {
    fmt.Println(fmt.Sprintf("%v: no disk device specified", nameUtility))
}

func main() {
    deviceName := ""
    if len(os.Args) < 2 {
        usage()
        os.Exit(1)
    } else {
        deviceName = os.Args[1]
    }
    greet()
    fmt.Println(fmt.Sprintf("\nYou choose device - %v", deviceName))
}
