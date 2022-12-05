package main

// import "github.com/Kaiser784/Proteus/config"
import (
	"fmt"
	"os"

	"github.com/Kaiser784/Proteus/config"
	// "strings"
	// "strconv"
)

func main() {
    // Check if at least one file name was provided as an argument
    if len(os.Args) < 2 {
        fmt.Println("Please provide at least one file name as an argument.")
        return
    }
	if len(os.Args) < 3 {
        fmt.Println("Please provide at least 2 files to generate a polyglot")
        return
    }

	fmt.Println(holder.Print())

    // Check if each file exists
    for _, fileName := range os.Args[1:] {
        if _, err := os.Stat(fileName); os.IsNotExist(err) {
            fmt.Printf("File not found: %s\n", fileName)
        } else {
            fmt.Printf("File found: %s\n", fileName)
        }
    }
}