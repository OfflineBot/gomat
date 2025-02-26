
package matrix3

import (
    "github.com/OfflineBot/gomat/matrix1"
    "github.com/OfflineBot/gomat/matrix2"
)

func (m *Matrix3[T]) checkShapes1(size1 int) {
    if m.Shape()[0] != size1 { panic("Shapes dont match!") }
}

func (m *Matrix3[T]) checkShapes2(size1, size2 int) {
    if m.Shape()[0] != size1 || m.Shape()[1] != size2 { panic("Shapes dont match!") }
}

func (m *Matrix3[T]) checkShapes3(size1, size2, size3 int) {
    if m.Shape()[0] != size1 || m.Shape()[1] != size2 || m.Shape()[2] != size3 { panic("Shapes dont match!") }
}

func (m *Matrix3[T]) checkInsideShape3(size1, size2, size3 int) {
    if (size1 < 1 || m.Shape()[0] >= size1) ||
        (size2 < 1 || m.Shape()[1] >= size2) ||
        (size3 < 1 || m.Shape()[2] >= size3) {
            panic("Shapes dont match!")
        }
}

// Add Matrix3 Elementwise
func (m *Matrix3[T]) AddMatrix3(other *Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val + other.GetValue(i, j, k), i, j, k)
            }
        }
    }
}


// Add Matrix2 Elementwise
func (m *Matrix3[T]) AddMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val + other.GetValue(i, j), i, j, k)
            }
        }
    }
}


// Add Matrix1 Elementwise
func (m *Matrix3[T]) AddMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val + other.GetValue(i), i, j, k)
            }
        }
    }
}


// Add Value at index
func (m *Matrix3[T]) AddIndex(value T, index1, index2, index3 int) {
    m.checkInsideShape3(index1, index2, index3)
    val := m.GetValue(index1, index2, index3)
    m.SetValue(val + value, index1, index2, index3)
}


// Sub Matrix3 Elementwise
func (m *Matrix3[T]) SubMatrix3(other *Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val - other.GetValue(i, j, k), i, j, k)
            }
        }
    }
}


// Sub Matrix2 Elementwise
func (m *Matrix3[T]) SubMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val - other.GetValue(i, j), i, j, k)
            }
        }
    }
}


// Sub Matrix1 Elementwise
func (m *Matrix3[T]) SubMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val - other.GetValue(i), i, j, k)
            }
        }
    }
}


// Sub Value at index
func (m *Matrix3[T]) SubIndex(value T, index1, index2, index3 int) {
    m.checkInsideShape3(index1, index2, index3)
    val := m.GetValue(index1, index2, index3)
    m.SetValue(val - value, index1, index2, index3)
}


// Mul Matrix3 Elementwise
func (m *Matrix3[T]) MulMatrix3(other *Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val * other.GetValue(i, j, k), i, j, k)
            }
        }
    }
}


// Mul Matrix2 Elementwise
func (m *Matrix3[T]) MulMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val * other.GetValue(i, j), i, j, k)
            }
        }
    }
}


// Mul Matrix1 Elementwise
func (m *Matrix3[T]) MulMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val * other.GetValue(i), i, j, k)
            }
        }
    }
}


// Mul Value at index
func (m *Matrix3[T]) MulIndex(value T, index1, index2, index3 int) {
    m.checkInsideShape3(index1, index2, index3)
    val := m.GetValue(index1, index2, index3)
    m.SetValue(val * value, index1, index2, index3)
}


// Div Matrix3 Elementwise
func (m *Matrix3[T]) DivMatrix3(other *Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val / other.GetValue(i, j, k), i, j, k)
            }
        }
    }
}


// Div Matrix2 Elementwise
func (m *Matrix3[T]) DivMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val / other.GetValue(i, j), i, j, k)
            }
        }
    }
}


// Div Matrix1 Elementwise
func (m *Matrix3[T]) DivMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := m.GetValue(i, j, k)
                m.SetValue(val / other.GetValue(i), i, j, k)
            }
        }
    }
}


// Div Value at index
func (m *Matrix3[T]) DivIndex(value T, index1, index2, index3 int) {
    m.checkInsideShape3(index1, index2, index3)
    val := m.GetValue(index1, index2, index3)
    m.SetValue(val / value, index1, index2, index3)
}


