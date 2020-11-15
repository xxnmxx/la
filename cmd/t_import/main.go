package main

import (
	"fmt"

	"github.com/xxnmxx/la"
)

func main() {
	mat := la.ImportCsv("csv.txt")
	fmt.Println(mat)
}
