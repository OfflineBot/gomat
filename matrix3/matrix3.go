

package matrix3

type Number interface {
    ~float32 | ~float64 | ~int
}


type Matrix3[T Number] struct {
    data [][][]T
    size1, size2, size3 int
}

