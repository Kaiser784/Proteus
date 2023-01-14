package config

import "fmt"

type AbstractFtype struct {
	DESC  string
	TYPE  string
	MAGIC []byte

	data       string
	cut        string
	prewrap    int64
	postwrap   int64
	bAppData   bool
	bParasite  bool
	parasite_o int64
	parasite_s int64
}

// constructor for abstract ftype
func NewFtype(DESC string, TYPE string, MAGIC []byte, data string) *AbstractFtype {
	return &AbstractFtype{
		DESC:  DESC,
		TYPE:  TYPE,
		MAGIC: MAGIC,
		data:  data,
	}
}

func (ftype *AbstractFtype) Identify() bool {
	return true
}

func (ftype *AbstractFtype) ShowDetails() {
	fmt.Println(ftype)
}
