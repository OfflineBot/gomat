

package matrix4

type Number interface {
    ~float32 | ~float64 | ~int
}


type Matrix4[T Number] struct {
    data [][][][]T
    size1, size2, size3, size4 int
}

