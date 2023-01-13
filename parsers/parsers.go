package parsers

import (
	"fmt"
	"github.com/Kaiser784/Proteus/config"
)

func JPG(ftype class.Ftype) {

	ftype.DESC = "JFIF / JPEG File Interchange Format"
	ftype.TYPE = "JPG"
	ftype.MAGIC = "FFD8"

	fmt.Println(ftype)
}

func PDF(ftype class.Ftype) {

	ftype.DESC = "Portable Document Format"
	ftype.TYPE = "PDF"
	ftype.MAGIC = "%PDF-1"

	fmt.Println(ftype)
}