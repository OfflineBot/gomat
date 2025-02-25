
package matrix4

func NewMatrix4[T Number](size1, size2, size3, size4 int) Matrix4[T] {
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

    return Matrix4[T]{
        data: data,
        size1: size1,
        size2: size2,
        size3: size3,
        size4: size4,
    }
}

