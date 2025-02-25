
package matrix3

func NewMatrix3[T Number](size1, size2, size3 int) Matrix3[T] {
    data := make([][][]T, size1)
    for i := range data {
        data[i] = make([][]T, size2)
        for j := range data[i] {
            data[i][j] = make([]T, size3)
        }
    }

    return Matrix3[T]{
        data: data,
        size1: size1,
        size2: size2,
        size3: size3,
    }
}

