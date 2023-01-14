package parsers

import (
	class "github.com/Kaiser784/Proteus/config"
)

type FtypeWrapper interface {
	Identify() bool
	ShowDetails()
}

type PDF struct {
	Ftype FtypeWrapper
}

// constructor for pdf
func NewPdf(data string) *PDF {
	const (
		DESC  = "Portable Document Format"
		TYPE  = "PDF"
		MAGIC = "%PDF-1"
	)

	magic := []byte(MAGIC)
	f := *class.NewFtype(DESC, TYPE, magic, data)

	return &PDF{
		//refer to https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
		Ftype: &f,
	}
}
