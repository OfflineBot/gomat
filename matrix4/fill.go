
package matrix4

import (
    "math/rand"
)

func (m *Matrix4[T]) FillZeros() {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    m.SetValue(0, i, j, k, h)
                }
            }
        }
    }
}


func (m *Matrix4[T]) FillRandom(min, max T) {
    for i := range m.Shape()[0] {
        for j := range m.Shape()[1] {
            for k := range m.Shape()[2] {
                for h := range m.Shape()[3] {
                    val := T(float64(min) + rand.Float64()*(float64(max)-float64(min)))
                    m.SetValue(val, i, j, k, h)
                }
            }
        }
    }
}

