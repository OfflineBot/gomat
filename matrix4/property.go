
package matrix4


func (m *Matrix4[T]) isValidIndex(index1, index2, index3, index4 int) bool {
    if (index1 < 0 || index1 >= m.size1) && 
        (index2 < 0 || index2 >= m.size2) &&
        (index3 < 0 || index3 >= m.size3) &&
        (index4 < 0 || index4 >= m.size4) {
        panic("Index not in range!")
    }
    return true
}



func (m *Matrix4[T]) Shape() [4]int {
    return [4]int {m.size1, m.size2, m.size3, m.size4}
}


func (m *Matrix4[T]) GetValue(index1, index2, index3, index4 int) T {
    m.isValidIndex(index1, index2, index3, index4)
    return m.data[index1][index2][index3][index4]
}


func (m *Matrix4[T]) SetValue(value T, index1, index2, index3, index4 int) {
    m.isValidIndex(index1, index2, index3, index4)
    m.data[index1][index2][index3][index4] = value
}



