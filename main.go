package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Kaiser784/Proteus/config"
    "github.com/Kaiser784/Proteus/config/parsers"
	"github.com/urfave/cli/v2"
	// "strings"
	// "strconv"
)

var outp string
var inp []string

func getExt() {
    for _,file := range inp {
        fmt.Println(file, "-", filepath.Ext(file))
    }
}

func check() {
    var ftype class.Ftype
    jpg.parse(ftype)
}

func main() {
    app := &cli.App{
		Name:  "Proteus",
        Usage: "Generates a file polyglot from the given input files",
        Flags: []cli.Flag{
            &cli.StringSliceFlag{
                Name:    "input",
                Aliases: []string{"i"},
                Usage:   "Input `file.txt, file.pdf`",
				Required: true,
            },
            &cli.StringFlag{
                Name:    "output",
                Aliases: []string{"o"},
                Usage:   "Output `filename`",
				Destination: &outp,
				Required: true,
            },
        },
        Action: func(cli *cli.Context) error {
            // fmt.Println("URL - ", host)
            fmt.Println("Output File:", outp);
            inp = cli.StringSlice("input")
            fmt.Println(inp)

            // generate()
            check()

            return nil
        },
        HideHelpCommand: true,
    }


    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }

    
}