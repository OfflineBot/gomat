
package matrix2

type Number interface {
    ~float32 | ~float64
}


type Matrix2[T Number] struct {
    data [][]T
    rows, cols int
}

