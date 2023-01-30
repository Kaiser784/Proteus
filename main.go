package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

func writeFile(data []byte, ext string) {

	name := "polyglot."+ext
	err := ioutil.WriteFile(name, data, 0644)
	if err != nil {
        panic(err)
    }
}

func stack(f1 *class.AbstractFtype, f2 *class.AbstractFtype, fn1 string, fn2 string) {
	fmt.Println("Stack : concatenation of File1 type", f1.TYPE,"and File2 type", f2.TYPE)
	data := append(f1.DATA, f2.DATA...)
	fmt.Println(fn1)
	fmt.Println(fn2)

	ext := f1.TYPE + "." +f2.TYPE
	writeFile(data, ext)
}

func check() {

	fdata1 := getByteFile(inp[0])
	fdata2 := getByteFile(inp[1])

	var ftype1 *class.AbstractFtype
	var ftype2 *class.AbstractFtype

	for _, ext := range PARSERS {
		f1, _ := getFtypeObject(ext, fdata1)
		if f1.Identify() {
			ftype1 = f1
		}
		f2, _ := getFtypeObject(ext, fdata2)
		if f2.Identify() {
			ftype2 = f2
		}
	}
	fmt.Println("ftype-1 - ", ftype1.TYPE)
	fmt.Println("ftype-2 - ", ftype2.TYPE)

	stack(ftype1, ftype2, inp[0], inp[1])
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
