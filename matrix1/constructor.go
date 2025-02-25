
package matrix1

func EmptyMatrix1[T Number](size int) *Matrix1[T] {

    if size < 1 {
        panic("Minimum length for Matrix is 1")
    }

    return &Matrix1[T]{
        data: make([]T, size), 
        size: size,
    }
}


func NewMatrix1[T Number](value []T) *Matrix1[T] {
    size := len(value)
    if size < 1 {
        panic("Minimum length for Matrix is 1")
    }
    return &Matrix1[T]{
        data: value,
        size: size,
    }
}


func RandomMatrix1[T Number](size int, min, max T) *Matrix1[T] {
    m := EmptyMatrix1[T](size)
    m.FillRandom(min, max)
    return m
}

