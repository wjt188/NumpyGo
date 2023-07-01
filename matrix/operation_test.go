package matrix

import "testing"

func TestRowVector_GetShape(t *testing.T) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3})
	t.Log(rv1.GetShape())
}

func TestRowVector_Add(t *testing.T) {
	rv1 := RowVector([]float64{1, 2, 3})
	rv2 := RowVector([]float64{1, 3, 5})
	t.Log(rv1.Add(rv2))
}
