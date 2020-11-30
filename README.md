# NoFileOpen

`NoFileOpen` find wrong usage of os.FileOpen

```go

package main

import (
    "fmt"
    "log"
	"os"
)

func main() { 
    file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666) 
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    fmt.Fprintln(file, "write")
}

```

```sh
$ go vet -vettool=`which NoFileOpen` main.go
./main.go:10:15: don't use os.OpenFile. Use os.Create
```
