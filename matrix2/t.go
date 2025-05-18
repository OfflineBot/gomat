package matrix2


func (m* Matrix2[T]) T() Matrix2[T] {
	out := EmptyMatrix2[T](m.Shape()[1], m.Shape()[0])

	for i := range m.Shape()[0] {
		for j := range m.Shape()[1] {
			val := m.GetValue(i, j)
			out.SetValue(val, j, i)
		}
	}

	return *out
}
