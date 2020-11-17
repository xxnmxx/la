package main

import (
	"github.com/xxnmxx/la"
	"os"
)

func main() {
	a := la.ImportCsv("a.txt")
	b := la.ImportCsv("b.txt")
	dot := la.Dot(a,b)
	dot.Output(os.Stdout)
}
