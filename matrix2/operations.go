
package matrix2

import "github.com/OfflineBot/gomat/matrix1"

func (m *Matrix2[T]) checkShapes1(size1 int) {
    if m.Shape()[0] != size1 { panic("Shapes dont match!") }
}

func (m *Matrix2[T]) checkShapes2(size1, size2 int) {
    if m.Shape()[0] != size1 || m.Shape()[1] != size2 { panic("Shapes dont match!") }
}

func (m *Matrix2[T]) checkInsideShape2(size1, size2 int) {
    if (size1 < 1 || m.Shape()[0] >= size1) ||
        (size2 < 1 || m.Shape()[1] >= size2) {
            panic("Shapes dont match!")
        }
}


// Add Matrix2 Elementwise
func (m *Matrix2[T]) AddMatrix2(other *Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val + other.GetValue(i, j), i, j)
        }
    }
}


// Add Matrix1 Elementwise
func (m *Matrix2[T]) AddMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val + other.GetValue(j), i, j)
        }
    }
}


// Add Value at index
func (m *Matrix2[T]) AddIndex(value T, index1, index2 int) {
    m.checkInsideShape2(index1, index2)
    val := m.GetValue(index1, index2)
    m.SetValue(val + value, index1, index2)
}


// Sub Matrix2 Elementwise
func (m *Matrix2[T]) SubMatrix2(other *Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val - other.GetValue(i, j), i, j)
        }
    }
}


// Sub Matrix1 Elementwise
func (m *Matrix2[T]) SubMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val - other.GetValue(i), i, j)
        }
    }
}


// Sub Value at index
func (m *Matrix2[T]) SubIndex(value T, index1, index2 int) {
    m.checkInsideShape2(index1, index2)
    val := m.GetValue(index1, index2)
    m.SetValue(val - value, index1, index2)
}


// Mul Matrix2 Elementwise
func (m *Matrix2[T]) MulMatrix2(other *Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val * other.GetValue(i, j), i, j)
        }
    }
}


// Mul Matrix1 Elementwise
func (m *Matrix2[T]) MulMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val * other.GetValue(i), i, j)
        }
    }
}


// Mul Value at index
func (m *Matrix2[T]) MulIndex(value T, index1, index2 int) {
    m.checkInsideShape2(index1, index2)
    val := m.GetValue(index1, index2)
    m.SetValue(val * value, index1, index2)
}


// Div Matrix2 Elementwise
func (m *Matrix2[T]) DivMatrix2(other *Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val / other.GetValue(i, j), i, j)
        }
    }
}


// Div Matrix1 Elementwise
func (m *Matrix2[T]) DivMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := m.GetValue(i, j)
            m.SetValue(val / other.GetValue(i), i, j)
        }
    }
}


// Div Value at index
func (m *Matrix2[T]) DivIndex(value T, index1, index2 int) {
    m.checkInsideShape2(index1, index2)
    val := m.GetValue(index1, index2)
    m.SetValue(val / value, index1, index2)
}


