
package matrix4

import (
    "github.com/OfflineBot/gomat/matrix1"
    "github.com/OfflineBot/gomat/matrix2"
    "github.com/OfflineBot/gomat/matrix3"
)

func (m *Matrix4[T]) checkShapes1(size1 int) {
    if m.Shape()[0] != size1 { panic("Shapes dont match!") }
}

func (m *Matrix4[T]) checkShapes2(size1, size2 int) {
    if m.Shape()[0] != size1 || m.Shape()[1] != size2 { panic("Shapes dont match!") }
}

func (m *Matrix4[T]) checkShapes3(size1, size2, size3 int) {
    if m.Shape()[0] != size1 || m.Shape()[1] != size2 || m.Shape()[2] != size3 { panic("Shapes dont match!") }
}

func (m *Matrix4[T]) checkShapes4(size1, size2, size3, size4 int) {
    if m.Shape()[0] != size1 || m.Shape()[1] != size2 || m.Shape()[2] != size3 || m.Shape()[3] != size4 { panic("Shapes dont match!") }
}


func (m *Matrix4[T]) checkInsideShape4(size1, size2, size3, size4 int) {
    if (size1 < 1 || m.Shape()[0] >= size1) ||
        (size2 < 1 || m.Shape()[1] >= size2) ||
        (size3 < 1 || m.Shape()[2] >= size3) ||
        (size4 < 1 || m.Shape()[3] >= size4) {
            panic("Shapes dont match!")
        }
}


// Add Matrix4 Elementwise
func (m *Matrix4[T]) AddMatrix4(other *Matrix4[T]) {
    m.checkShapes4(other.Shape()[0], other.Shape()[1], other.Shape()[2], other.Shape()[3])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val + other.GetValue(i, j, k, h), i, j, k, h)
                }
            }
        }
    }
}

// Add Matrix3 Elementwise
func (m *Matrix4[T]) AddMatrix3(other *matrix3.Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val + other.GetValue(i, j, k), i, j, k, h)
                }
            }
        }
    }
}


// Add Matrix2 Elementwise
func (m *Matrix4[T]) AddMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val + other.GetValue(i, j), i, j, k, h)
                }
            }
        }
    }
}


// Add Matrix1 Elementwise
func (m *Matrix4[T]) AddMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val + other.GetValue(i), i, j, k, h)
                }
            }
        }
    }
}


// Add Value at index
func (m *Matrix4[T]) AddIndex(value T, index1, index2, index3, index4 int) {
    m.checkInsideShape4(index1, index2, index3, index4)
    val := m.GetValue(index1, index2, index3, index4)
    m.SetValue(val + value, index1, index2, index3, index4)
}


// Sub Matrix4 Elementwise
func (m *Matrix4[T]) SubMatrix4(other *Matrix4[T]) {
    m.checkShapes4(other.Shape()[0], other.Shape()[1], other.Shape()[2], other.Shape()[3])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val - other.GetValue(i, j, k, h), i, j, k, h)
                }
            }
        }
    }
}

// Sub Matrix3 Elementwise
func (m *Matrix4[T]) SubMatrix3(other *matrix3.Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val - other.GetValue(i, j, k), i, j, k, h)
                }
            }
        }
    }
}


// Sub Matrix2 Elementwise
func (m *Matrix4[T]) SubMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val - other.GetValue(i, j), i, j, k, h)
                }
            }
        }
    }
}


// Sub Matrix1 Elementwise
func (m *Matrix4[T]) SubMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val - other.GetValue(i), i, j, k, h)
                }
            }
        }
    }
}


// Sub Value at index
func (m *Matrix4[T]) SubIndex(value T, index1, index2, index3, index4 int) {
    m.checkInsideShape4(index1, index2, index3, index4)
    val := m.GetValue(index1, index2, index3, index4)
    m.SetValue(val - value, index1, index2, index3, index4)
}


// Mul Matrix4 Elementwise
func (m *Matrix4[T]) MulMatrix4(other *Matrix4[T]) {
    m.checkShapes4(other.Shape()[0], other.Shape()[1], other.Shape()[2], other.Shape()[3])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val * other.GetValue(i, j, k, h), i, j, k, h)
                }
            }
        }
    }
}

// Mul Matrix3 Elementwise
func (m *Matrix4[T]) MulMatrix3(other *matrix3.Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val * other.GetValue(i, j, k), i, j, k, h)
                }
            }
        }
    }
}


// Mul Matrix2 Elementwise
func (m *Matrix4[T]) MulMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val * other.GetValue(i, j), i, j, k, h)
                }
            }
        }
    }
}


// Mul Matrix1 Elementwise
func (m *Matrix4[T]) MulMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val * other.GetValue(i), i, j, k, h)
                }
            }
        }
    }
}


// Mul Value at index
func (m *Matrix4[T]) MulIndex(value T, index1, index2, index3, index4 int) {
    m.checkInsideShape4(index1, index2, index3, index4)
    val := m.GetValue(index1, index2, index3, index4)
    m.SetValue(val * value, index1, index2, index3, index4)
}


// Div Matrix4 Elementwise
func (m *Matrix4[T]) DivMatrix4(other *Matrix4[T]) {
    m.checkShapes4(other.Shape()[0], other.Shape()[1], other.Shape()[2], other.Shape()[3])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val / other.GetValue(i, j, k, h), i, j, k, h)
                }
            }
        }
    }
}

// Div Matrix3 Elementwise
func (m *Matrix4[T]) DivMatrix3(other *matrix3.Matrix3[T]) {
    m.checkShapes3(other.Shape()[0], other.Shape()[1], other.Shape()[2])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val / other.GetValue(i, j, k), i, j, k, h)
                }
            }
        }
    }
}


// Div Matrix2 Elementwise
func (m *Matrix4[T]) DivMatrix2(other *matrix2.Matrix2[T]) {
    m.checkShapes2(other.Shape()[0], other.Shape()[1])
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val / other.GetValue(i, j), i, j, k, h)
                }
            }
        }
    }
}


// Sub Matrix1 Elementwise
func (m *Matrix4[T]) DivMatrix1(other *matrix1.Matrix1[T]) {
    m.checkShapes1(other.Shape())
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := m.GetValue(i, j, k, h)
                    m.SetValue(val / other.GetValue(i), i, j, k, h)
                }
            }
        }
    }
}


// Sub Value at index
func (m *Matrix4[T]) DivIndex(value T, index1, index2, index3, index4 int) {
    m.checkInsideShape4(index1, index2, index3, index4)
    val := m.GetValue(index1, index2, index3, index4)
    m.SetValue(val / value, index1, index2, index3, index4)
}



