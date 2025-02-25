
package matrix4

func EmptyMatrix4[T Number](size1, size2, size3, size4 int) *Matrix4[T] {
    data := make([][][][]T, size1)
    for i := range data {
        data[i] = make([][][]T, size2)
        for j := range data[i] {
            data[i][j] = make([][]T, size3)
            for k := range data[i][j] {
                data[i][j][k] = make([]T, size4)
            }
        }
    }

    return &Matrix4[T]{
        data: data,
        size1: size1,
        size2: size2,
        size3: size3,
        size4: size4,
    }
}


func NewMatrix4[T Number](value [][][][]T) *Matrix4[T] {
    size1 := len(value)
    if size1 < 1 { panic("Minimum length for Matrix is 1") }
    size2 := len(value[0])
    if size2 < 1 { panic("Minimum length for Matrix is 1") }
    size3 := len(value[0][0])
    if size3 < 1 { panic("Minimum length for Matrix is 1") }
    size4 := len(value[0][0][0])
    if size4 < 1 { panic("Minimum length for Matrix is 1") }

    return &Matrix4[T]{
        data: value,
        size1: size1,
        size2: size2,
        size3: size3,
        size4: size4,
    }
}


func RandomMatrix1[T Number](size1, size2, size3, size4 int, min, max T) *Matrix4[T] {
    m := EmptyMatrix4[T](size1, size2, size3, size4)
    m.FillRandom(min, max)
    return m
}

