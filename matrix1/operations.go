
package matrix1


func (m *Matrix1[T]) checkShapes(size int) {
    if m.Shape() != size { panic("Shapes dont match!") }
}

// Add Matrix Elementwise
func (m *Matrix1[T]) AddMatrix1(other *Matrix1[T]) {
    m.checkShapes(other.Shape())
    for i := range m.Shape() {
        val := m.GetValue(i)
        m.SetValue(val + other.GetValue(i), i)
    }
}


// Add Value at index
func (m *Matrix1[T]) AddIndex(value T, index int) {
    m.checkShapes(index)
    val := m.GetValue(index)
    m.SetValue(val + value, index)
}


// Sub Matrix Elementwise
func (m *Matrix1[T]) SubMatrix1(other *Matrix1[T]) {
    m.checkShapes(other.Shape())
    for i := range m.Shape() {
        val := m.GetValue(i)
        m.SetValue(val - other.GetValue(i), i)
    }
}


// Sub Value at index
func (m *Matrix1[T]) SubIndex(value T, index int) {
    m.checkShapes(index)
    val := m.GetValue(index)
    m.SetValue(val - value, index)
}


// Mul Matrix Elementwise
func (m *Matrix1[T]) MulMatrix1(other *Matrix1[T]) {
    m.checkShapes(other.Shape())
    for i := range m.Shape() {
        val := m.GetValue(i)
        m.SetValue(val * other.GetValue(i), i)
    }
}


// Mul Value at index
func (m *Matrix1[T]) MulIndex(value T, index int) {
    m.checkShapes(index)
    val := m.GetValue(index)
    m.SetValue(val * value, index)
}


// Div Matrix Elementwise
func (m *Matrix1[T]) DivMatrix1(other *Matrix1[T]) {
    m.checkShapes(other.Shape())
    for i := range m.Shape() {
        val := m.GetValue(i)
        m.SetValue(val / other.GetValue(i), i)
    }
}


// Div Value at index
func (m *Matrix1[T]) DivIndex(value T, index int) {
    m.checkShapes(index)
    val := m.GetValue(index)
    m.SetValue(val / value, index)
}

