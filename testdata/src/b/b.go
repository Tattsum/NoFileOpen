package b

import os2 "os"

func main() {
	_, _ = os2.OpenFile("fileName", os2.O_RDONLY, 0666) // NG
	_, _ = os2.Create("fileName") //OK

	fn := "fileName"
	flag := os2.O_RDONLY
	_, _ = os2.OpenFile(fn, flag, 0666) // NG
}
