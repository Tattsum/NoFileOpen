package a

import (
	"os"
	os2 "os"
)

func f() {
	_, _ = os.OpenFile("fileName", os.O_CREATE,0666) // want "NG"
	_, _ = os.Create("fileName") // OK

	_, _ = os2.OpenFile("fileName", os.O_CREATE,0666) // want "NG"
	_, _ = os2.Create("fileName") // OK

}
