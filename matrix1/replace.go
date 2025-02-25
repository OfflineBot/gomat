
package matrix1


func (m *Matrix1[T]) ReplaceAllBy(before, after T) {
    for i := range m.Shape() {
        if m.GetValue(i) == before {
            m.SetValue(after, i)
        }
    }
}
