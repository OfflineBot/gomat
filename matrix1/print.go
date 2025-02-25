
package matrix1

import "fmt"

func (m *Matrix1[T]) Println() {
    fmt.Print("{ ")
    fmt.Print(m.GetValue(0))
    for i := 1; i < m.Shape(); i++ {
        fmt.Print(", ")
        fmt.Print(m.GetValue(i))
    }
    fmt.Printf(" } [Shape: %d] \n", m.Shape())
}

