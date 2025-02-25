
package matrix2

import "fmt"

func (m *Matrix2[T]) Println() {
    fmt.Print("{ \n")
    for i := range m.Shape()[0] {
        fmt.Print("  { ")
        for j := range m.Shape()[1] - 1 {
            fmt.Print(m.GetValue(i, j))
            fmt.Print(", ")
        }
        fmt.Print(m.GetValue(i, m.Shape()[1]-1))
        fmt.Print(" }, \n")
    }
    fmt.Printf("} [Shape: %dx%d] \n", m.Shape()[0], m.Shape()[1])
}

