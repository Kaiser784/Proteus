package main

// import "github.com/Kaiser784/Proteus/config"
import (
	"fmt"
	"os"
	"github.com/Kaiser784/Proteus/config"
    "flag"
	// "strings"
	// "strconv"
)


func main() {

    var inputFlag StringArray
	flag.Var(&inputFlag, "input", "the input files")
    outputFlag := flag.String("output", "", "the output file")
    flag.Parse()
    
    if *outputFlag == "" {
		// If not, throw an error
		fmt.Println("Error: missing output flag")
		return
	}

    if len(inputFlag) < 2 {
        fmt.Println("Input atleast 2 files to create a polyglot")
        fmt.Println("input the files separately")
        fmt.Println("--input file1 --input file2")
        return
    }

	holder.Print("Module is working")

    fmt.Println("Output file:", *outputFlag)

    // Check if each file exists
    for _, fileName := range inputFlag {
        if _, err := os.Stat(fileName); os.IsNotExist(err) {
            fmt.Printf("Input File not found: %s\n", fileName)
        } else {
            fmt.Printf("Input File found: %s\n", fileName)
        }
    }
}

// StringArray is a custom type that satisfies the flag.Value interface
type StringArray []string

// String returns a string representation of the StringArray
func (s *StringArray) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set adds a new string to the StringArray
func (s *StringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}