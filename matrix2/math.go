
package matrix2


func (m *Matrix2[T]) Dot(other *Matrix2[T]) *Matrix2[T] {
    shape1 := m.Shape()[0]
    shape2 := m.Shape()[1]
    shape3 := other.Shape()[0]
    shape4 := other.Shape()[1]

    if shape2 != shape3 {
        panic("Shapes dont match for Matrix multiplication")
    }

    out := EmptyMatrix2[T](shape1, shape4)

    for i := range shape1 {
        for j := range shape2 {
            for k := range shape4 {
                val := out.GetValue(i, k)
                val2 := val + (m.GetValue(i, j) * other.GetValue(j, k))
                out.SetValue(val2, i, k)
            }
        }
    }

    return out
}


