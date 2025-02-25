package matrix1

import (
	"math/rand"
)

func (m *Matrix1[T]) FillZeros() {
    for i := range m.Shape() {
        m.SetValue(0, i)
    }
}


func (m *Matrix1[T]) FillRandom(min, max T) {
    for i := range m.Shape() {
        val := T(float64(min) + rand.Float64()*(float64(max)-float64(min)))
        m.SetValue(val, i)
    }
}

