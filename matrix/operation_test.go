package matrix

import "testing"

func TestRowVector_GetShape(t *testing.T) {
	rv1 := RowVector([]float64{1.1, 2.2, 3.3})
	t.Log(rv1.GetShape())
}
