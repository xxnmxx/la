package la

import "log"

func Dot(a,b *Matrix) *Matrix {
	if a.Cols != b.Rows {
		log.Fatalf("Shape unmatch: a(%v, %v), b(%v, %v)", a.Rows, a.Cols, b.Rows, b.Cols)
	}
	d := make([]float64, a.Rows * b.Cols)
	for i := range d {
			for k := 0; k < b.Rows; k++ {
				d[i] += a.Data[(i/b.Cols)*a.Cols + k] * b.Data[k*b.Cols + (i%b.Cols)]
			}
	}
	return &Matrix{
		Data: d,
		Rows: a.Rows,
		Cols: b.Cols,
	}
}

