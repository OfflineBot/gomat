
package matrix4

import "fmt"

func (m *Matrix4[T]) Println() {
    fmt.Print("{ \n")

    for i := range m.Shape()[0] {
        fmt.Print("  { \n") 
        for j := range m.Shape()[1] {
            fmt.Print("    { \n") 
            for k := range m.Shape()[2] {
                fmt.Print("      { ")
                for h := range m.Shape()[3] - 1 {
                    fmt.Print(m.GetValue(i, j, k, h))
                    fmt.Print(", ")
                }
                fmt.Print(m.GetValue(i, j, k, m.Shape()[3]-1))
                fmt.Print(" }, \n")
            }
            fmt.Print("    },\n")
        }
        fmt.Print("  }, \n")
    }

    fmt.Printf("} [Shape: %dx%dx%dx%d] \n", m.Shape()[0], m.Shape()[1], m.Shape()[2], m.Shape()[3])
}

