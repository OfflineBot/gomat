
package matrix2

func NewMatrix2[T Number](rows, cols int) Matrix2[T] {
    data := make([][]T, rows)
    for i := range data {
        data[i] = make([]T, cols)
    }

    return Matrix2[T]{
        data: data,
        rows: rows,
        cols: cols,
    }
}

