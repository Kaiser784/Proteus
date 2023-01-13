package jpg

import (
	"fmt"
	"github.com/Kaiser784/Proteus/config"
)

func parse(ftype class.Ftype) {

	ftype.DESC = "JFIF / JPEG File Interchange Format"
	ftype.TYPE = "JPG"
	ftype.MAGIC = "FFD8"

	class.class()
	fmt.Println(ftype)
}