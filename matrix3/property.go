
package matrix3


func (m *Matrix3[T]) isValidIndex(index1, index2, index3 int) bool {
    if (index1 < 0 || index1 >= m.size1) ||
        (index2 < 0 || index2 >= m.size2) ||
        (index3 < 0 || index3 >= m.size3) {
        panic("Index not in range!")
    }
    return true
}



func (m *Matrix3[T]) Shape() [3]int {
    return [3]int {m.size1, m.size2, m.size3}
}


func (m *Matrix3[T]) GetValue(index1, index2, index3 int) T {
    m.isValidIndex(index1, index2, index3)
    return m.data[index1][index2][index3]
}


func (m *Matrix3[T]) SetValue(value T, index1, index2, index3 int) {
    m.isValidIndex(index1, index2, index3)
    m.data[index1][index2][index3] = value
}



