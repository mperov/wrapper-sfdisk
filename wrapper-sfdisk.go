package main


import (
    "fmt"
    "os"
)

const nameUtility = "wrapper-disk"

func greet() {
    fmt.Println(fmt.Sprintf("Welcome to %v!", nameUtility))
    fmt.Println("Changes which is written on disk can destroy dates on your disk!")
    fmt.Println("Please, be careful using this utility!")
}

func usage() {
    fmt.Println(fmt.Sprintf("%v: no disk device specified", nameUtility))
}

func main() {
    //device := ""
    if len(os.Args) < 2 {
        usage()
        os.Exit(1)
    } else {
        //device = os.Args[1]
    }
    greet()
}
