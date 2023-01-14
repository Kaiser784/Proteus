package parsers

import (
	class "github.com/Kaiser784/Proteus/config"
)

type PDF struct {
	Ftype class.AbstractFtype
}

// constructor for pdf
func NewPdf(data string) *PDF {
	const (
		DESC  = "Portable Document Format"
		TYPE  = "PDF"
		MAGIC = "%PDF-1"
	)

	magic := []byte(MAGIC)

	return &PDF{
		Ftype: *class.NewFtype(DESC, TYPE, magic, data),
	}
}
