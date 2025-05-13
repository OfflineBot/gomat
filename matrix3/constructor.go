
package matrix3


func EmptyMatrix3[T Number](size1, size2, size3 int) *Matrix3[T] {
    data := make([][][]T, size1)
    for i := range data {
        data[i] = make([][]T, size2)
        for j := range data[i] {
            data[i][j] = make([]T, size3)
        }
    }

    return &Matrix3[T]{
        data: data,
        size1: size1,
        size2: size2,
        size3: size3,
    }
}


func NewMatrix3[T Number](value [][][]T) *Matrix3[T] {
    size1 := len(value)
    if size1 < 1 { panic("Minimum length for Matrix is 1") }
    size2 := len(value[0])
    if size2 < 1 { panic("Minimum length for Matrix is 1") }
    size3 := len(value[0][0])
    if size3 < 1 { panic("Minimum length for Matrix is 1") }

    return &Matrix3[T]{
        data: value,
        size1: size1,
        size2: size2,
        size3: size3,
    }
}


func RandomMatrix3[T Number](size1, size2, size3 int, min, max T) *Matrix3[T] {
    m := EmptyMatrix3[T](size1, size2, size3)
    m.FillRandom(min, max)
    return m
}

