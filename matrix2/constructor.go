
package matrix2

func EmptyMatrix2[T Number](rows, cols int) *Matrix2[T] {
    data := make([][]T, rows)
    for i := range data {
        data[i] = make([]T, cols)
    }

    return &Matrix2[T]{
        data: data,
        rows: rows,
        cols: cols,
    }
}


func NewMatrix2[T Number](value [][]T) *Matrix2[T] {
    size1 := len(value)
    if size1 < 1 { panic("Minimum length for Matrix is 1") }
    size2 := len(value[0])
    if size2 < 1 { panic("Minimum length for Matrix is 1") }
    return &Matrix2[T]{
        data: value,
        rows: size1,
        cols: size2,
    }
}


func RandomMatrix1[T Number](rows, cols int, min, max T) *Matrix2[T] {
    m := EmptyMatrix2[T](rows, cols)
    m.FillRandom(min, max)
    return m
}

