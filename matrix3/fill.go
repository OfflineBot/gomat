
package matrix3

import (
    "math/rand"
)

func (m *Matrix3[T]) FillZeros() {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                m.SetValue(0, i, j, k)
            }
        }
    }
}


func (m *Matrix3[T]) FillRandom(min, max T) {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                val := T(float64(min) + rand.Float64()*(float64(max)-float64(min)))
                m.SetValue(val, i, j, k)
            }
        }
    }
}

