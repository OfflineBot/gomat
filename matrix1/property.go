
package matrix1


func (m *Matrix1[T]) isValidIndex(index int) bool {
    if index < 0 || index >= m.size {
        panic("Index not in range!")
    }
    return true
}



func (m *Matrix1[T]) Shape() int {
    return m.size
}


func (m *Matrix1[T]) GetValue(index int) T {
    m.isValidIndex(index)
    return m.data[index]
}


func (m *Matrix1[T]) SetValue(value T, index int) {
    m.isValidIndex(index)
    m.data[index] = value
}



