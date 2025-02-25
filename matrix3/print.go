
package matrix3

import "fmt"

func (m *Matrix3[T]) Println() {
    fmt.Print("{ \n")

    for i := range m.Shape()[0] {
        fmt.Print("  { \n") 
        for j := range m.Shape()[1] {
            fmt.Print("    { ") 
            for k := range m.Shape()[2] - 1 {
                fmt.Print(m.GetValue(i, j, k))
                fmt.Print(", ")
            }
            fmt.Print(m.GetValue(i, j, m.Shape()[2]-1))
            fmt.Print(" }, \n")
        }
        fmt.Print("  }, \n")
    }

    fmt.Printf("} [Shape: %dx%dx%d] \n", m.Shape()[0], m.Shape()[1], m.Shape()[2])
}

