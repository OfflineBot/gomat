
package matrix2


func (m *Matrix2[T]) isValidIndex(index1, index2 int) bool {
    if (index1 < 0 || index1 >= m.rows) || (index2 < 0 || index2 >= m.cols) {
        panic("Index not in range!")
    }
    return true
}



func (m *Matrix2[T]) Shape() [2]int {
    return [2]int {m.rows, m.cols}
}


func (m *Matrix2[T]) GetValue(rows, cols int) T {
    m.isValidIndex(rows, cols)
    return m.data[rows][cols]
}


func (m *Matrix2[T]) SetValue(value T, index1, index2 int) {
    m.isValidIndex(index1, index2)
    m.data[index1][index2] = value
}



