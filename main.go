package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"bufio"
	"io"
	"errors"

	class "github.com/Kaiser784/Proteus/config"
	"github.com/Kaiser784/Proteus/parsers"
	"github.com/urfave/cli/v2"
	// "strings"
	// "strconv"
)

var outp string
var inp []string
var PARSERS = []string{"JPG", "PDF"}

func getExt() {
	for _, file := range inp {
		fmt.Println(file, "-", filepath.Ext(file))
	}
}

func getByteFile(filename string) []byte{
	f, _ := os.Open(filename)
	br := bufio.NewReader(f)
	fdata := []byte("")

	for {
		b,err := br.ReadByte()
		fdata = append(fdata, b)

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}
	
		if err != nil {
			// end of file
			break
		}
	}
	return fdata
}

func getFtypeObject(fileExtension string, data []byte) (*class.AbstractFtype, bool) {
	switch fileExtension {
	case "PDF":
		return parsers.NewPdf(data), false
	case "JPG":
		return parsers.NewJpg(data), false
	}
	return nil, true
}

func check() {

	fdata1 := getByteFile(inp[0])
	fdata2 := getByteFile(inp[1])

	var ftype1 string
	var ftype2 string

	for _, ext := range PARSERS {
		f1, _ := getFtypeObject(ext, fdata1)
		if f1.Identify() {
			ftype1 = ext
		}
		f2, _ := getFtypeObject(ext, fdata2)
		if f2.Identify() {
			ftype2 = ext
		}
	}
	fmt.Println("ftype-1 - ", ftype1)
	fmt.Println("ftype-2 - ", ftype2)
}

func main() {
	app := &cli.App{
		Name:  "Proteus",
		Usage: "Generates a file polyglot from the given input files",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Usage:    "Input `file.txt, file.pdf`",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Output `filename`",
				Destination: &outp,
				Required:    true,
			},
		},
		Action: func(cli *cli.Context) error {
			// fmt.Println("URL - ", host)
			fmt.Println("Output File:", outp)
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
