
package matrix3

func (m *Matrix3[T]) ReplaceAllBy(before, after T) {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                if m.GetValue(i, j, k) == before {
                    m.SetValue(after, i, j, k)
                }
            }
        }
    }
}

