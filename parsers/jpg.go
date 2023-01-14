package parsers

import (
	class "github.com/Kaiser784/Proteus/config"
)

type JPG struct {
	Ftype class.AbstractFtype
}

// constructor for jpg
func NewJpg(data string) *JPG {
	const (
		DESC  = "JFIF / JPEG File Interchange Format"
		TYPE  = "JPG"
		MAGIC = "\xFF\xD8"
	)

	magic := []byte(MAGIC)

	return &JPG{
		Ftype: *class.NewFtype(DESC, TYPE, magic, data),
	}
}
