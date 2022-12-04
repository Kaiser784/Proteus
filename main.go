package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

const (
    PDF = iota + 1
    ZIP
    PNG
)

var options = map[int]string{
    PDF: "pdf",
    ZIP: "zip",
    PNG: "png",
}

func main() {
    // Check if a file name was provided as an argument
    if len(os.Args) < 2 {
        fmt.Println("Please provide a file name as an argument.")
        return
    }
    fileName := os.Args[1]

    fmt.Println("Filename: ", fileName)

    var selectedOptions []int
    var input string
    fmt.Println("Please select one or more options (separated by commas without any spaces):")
    for option, message := range options {
        fmt.Printf("%d. %s\n", option, message)
    }
    fmt.Scanf("%s", &input)
    for _, option := range strings.Split(input, ",") {
        if option, err := strconv.Atoi(option); err == nil {
            if _, ok := options[option]; ok {
                selectedOptions = append(selectedOptions, option)
            }
        }
    }
    if len(selectedOptions) == 0 {
        fmt.Println("No options selected.")
    } else {
        fmt.Println("Selected options:")
        for _, option := range selectedOptions {
            fmt.Println(options[option])
        }
    }
}