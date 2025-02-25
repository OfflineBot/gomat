
package matrix1

func NewMatrix1[T Number](size int) Matrix1[T] {
    return Matrix1[T]{
        data: make([]T, size), 
        size: size,
    }
}

