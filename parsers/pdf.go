package parsers

import (
	class "github.com/Kaiser784/Proteus/config"
)

// constructor for pdf
func NewPdf(data []byte) *class.AbstractFtype {
	const (
		DESC  = "Portable Document Format"
		TYPE  = "PDF"
		MAGIC = "%PDF-1"
	)

	magic := []byte(MAGIC)

	return class.NewFtype(DESC, TYPE, magic, data)
}
