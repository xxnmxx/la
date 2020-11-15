package main

import (
	"fmt"
	"os"

	"github.com/xxnmxx/la"
)

func main() {
	mat := la.ImportCsv("csv.txt")
	fmt.Println(mat)
	mat.Output(os.Stdout)
}
