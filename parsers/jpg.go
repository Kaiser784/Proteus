package parsers

import (
	class "github.com/Kaiser784/Proteus/config"
)

type JPG struct {
	Ftype FtypeWrapper
}

func NewJpg(data string) *JPG {
	const (
		DESC  = "JFIF / JPEG File Interchange Format"
		TYPE  = "JPG"
		MAGIC = "\xFF\xD8"
	)

	magic := []byte(MAGIC)
	f := *class.NewFtype(DESC, TYPE, magic, data)

	return &JPG{
		Ftype: &f,
	}
}
