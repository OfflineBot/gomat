

package matrix1

type Number interface {
    ~float32 | ~float64
}


type Matrix1[T Number] struct {
    data []T
    size int
}

