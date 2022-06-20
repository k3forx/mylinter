package main

import (
	"github.com/k3forx/mylinter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mylinter.Analyzer) }
