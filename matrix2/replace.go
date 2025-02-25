
package matrix2


func (m *Matrix2[T]) ReplaceAllBy(before, after T) {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            if m.GetValue(i, j) == before {
                m.SetValue(after, i, j)
            }
        }
    }
}

