
package matrix4

func (m *Matrix4[T]) ReplaceAllBy(before, after T) {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    if m.GetValue(i, j, k, h) == before {
                        m.SetValue(after, i, j, k, h)
                    }
                }
            }
        }
    }
}

