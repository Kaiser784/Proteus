package config

import (
	"bytes"
	"fmt"
)

type AbstractFtype struct {
	DESC  string
	TYPE  string
	MAGIC []byte

	DATA       []byte
	cut        string
	prewrap    int64
	postwrap   int64
	bAppData   bool
	bParasite  bool
	parasite_o int64
	parasite_s int64
}

// constructor for abstract ftype
func NewFtype(DESC string, TYPE string, MAGIC []byte, data []byte) *AbstractFtype {
	return &AbstractFtype{
		DESC:  DESC,
		TYPE:  TYPE,
		MAGIC: MAGIC,
		DATA:  data,
	}
}

func (ftype *AbstractFtype) Identify() bool {
	return bytes.HasPrefix(ftype.DATA, ftype.MAGIC)
}

func (ftype *AbstractFtype) ShowDetails() {
	fmt.Println(ftype)
}
