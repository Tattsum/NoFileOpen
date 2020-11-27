package main

import (
	"github.com/Tattsum/NoFileOpen"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(NoFileOpen.Analyzer) }
