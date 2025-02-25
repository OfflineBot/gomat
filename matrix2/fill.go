
package matrix2

import (
    "math/rand"
)

func (m *Matrix2[T]) FillZeros() {
    for i := range m.Shape()[0] {
        for k := range m.Shape()[1] {
            m.SetValue(0, i, k)
        }
    }
}


func (m *Matrix2[T]) FillRandom(min, max T) {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            val := T(float64(min) + rand.Float64()*(float64(max)-float64(min)))
            m.SetValue(val, i, j)
        }
    }
}

