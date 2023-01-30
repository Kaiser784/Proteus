package parsers

import (
	class "github.com/Kaiser784/Proteus/config"
)

// constructor for jpg
func NewJpg(data []byte) *class.AbstractFtype {
	const (
		DESC  = "JFIF / JPEG File Interchange Format"
		TYPE  = "JPG"
		MAGIC = "\xFF\xD8"
	)

	magic := []byte(MAGIC)

	return class.NewFtype(DESC, TYPE, magic, data)
}
