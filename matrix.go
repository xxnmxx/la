package la

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Data []float64
	Rows int
	Cols int
}

func NewMat(d []float64, r, c int) *Matrix {
	return &Matrix{
		Data: d,
		Rows: r,
		Cols: c,
	}
}

func ImportCsv(p string) *Matrix {
	f, err := os.Open(p)
	if err != nil {
		log.Fatalf("file open error", err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	data := make([]float64, 0)
	rows := 0
	cols := 0
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("import error: %v", err)
		}
		// do this only in the first loop.
		if rows == 0 {
			cols = len(rec)
		}
		if cols != len(rec) {
			log.Fatal("import error: shape unmatch")
		}
		elms := convToFloat(rec)
		//for _, v := range elms {
		//	data = append(data, v)
		//}
		appendElms(elms, &data)
		rows++
	}
	return NewMat(data, rows, cols)
}

func convToFloat(s []string) []float64 {
	data := make([]float64, len(s))
	for i, v := range s {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatalf("import error: parse: %v %v)", i, v)
		}
		data[i] = f
	}
	return data
}

func appendElms(src []float64, dst *[]float64) {
	for _, v := range src {
		*dst = append(*dst, v)
	}
}

func (m *Matrix) Output(o os.Stdout) {
	var b strings.Builder
}
