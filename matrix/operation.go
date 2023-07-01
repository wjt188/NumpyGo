package matrix

import "errors"

// 向量运算接口
func (rv1 RowVector) GetShape() (s [2]int) {
	s[0] = 1
	s[1] = len(rv1)
	return
}
func (rv1 *RowVector) Mul(c float64) {
	l := rv1.GetShape()[1]
	rvTemp := make([]float64, l, l)
	copy(rvTemp, *rv1)
	for i := 0; i < l; i++ {
		rvTemp[i] *= c
	}
	copy(*rv1, rvTemp)
}

func NewRowVector(l int) (rv1 RowVector, err error) {
	if l <= 0 {
		err = errors.New("hte dimension of row vectors must > 0")
		return
	}
	rv1 = make([]float64, l, l)
	return
}
func (rv1 RowVector) Add(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimension of the vetcors are not equal!!")
		return
	}
	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, v := range rv1 {
		rv3[i] = v + rv2[i]
	}
	return
}
func (rv1 RowVector) Minus(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimension of the vetcors are not equal!!")
		return
	}
	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, v := range rv1 {
		rv3[i] = v - rv2[i]
	}
	return
}
func (rv1 RowVector) Dot(rv2 RowVector) (dot float64, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimension of the vetcors are not equal!!")
		return
	}
	for i, v := range rv1 {
		dot += v * rv2[i]
	}
	return
}
func (rv1 RowVector) Cross(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimension of the vetcors are not equal!!")
		return
	}
	dim := rv1.GetShape()[1]
	if dim != 2 && dim != 3 {
		err = errors.New("只能计算二维和三维的向量相乘")
	}
	rv3, _ = NewRowVector(3)
	switch dim {
	case 2:
		rv3[0] = 0
		rv3[1] = 0
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	case 3:
		rv3[0] = rv1[1]*rv2[1] - rv1[2]*rv2[1]
		rv3[1] = rv2[2]*rv2[0] - rv1[0]*rv2[2]
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]

	}
	return
}
func (rv1 RowVector) Transpose() (cv ColumVector) {
	cv = make([][1]float64, rv1.GetShape()[1], rv1.GetShape()[1])
	for i, v := range rv1 {
		cv[i][0] = v
	}
	return
}

// 矩阵运算接口
func (m1 Matrix) GetShape() (s [2]int) {
	s[0] = len(m1)
	s[1] = len(m1[0])
	return
}

func (m1 *Matrix) Mul(c float64) {
	row, col := m1.GetShape()[0], m1.GetShape()[1]
	mTemp := make([][]float64, row, col)
	copy(mTemp, *m1)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			mTemp[i][j] *= c
		}
		copy(*m1, mTemp)
	}
}
func NewMartix(r, c int) (mat Matrix, err error) {
	if r <= 0 || c <= 0 {
		err = errors.New("")
		return
	}
	mat = make([][]float64, r, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]float64, c, c)
	}
	return
}

func (m1 Matrix) Add(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two Matrix are not equal")
		return
	}
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMartix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return

}
func (m1 Matrix) Minus(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two Matrix are not equal")
		return
	}
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMartix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] - m2[i][j]
		}
	}
	return

}

func (m1 Matrix) MatMul(m2 Matrix) (m3 Matrix, err error) {
	r1, c1 := m1.GetShape()[0], m1.GetShape()[1]
	r2, c2 := m2.GetShape()[0], m2.GetShape()[1]
	if c1 != c2 {
		err = errors.New("两个矩阵不匹配")
		return
	}
	m3, _ = NewMartix(r1, c2)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			v1 := RowVector(m1[i])
			v2 := make([]float64, r2, r2)
			for k := 0; k < r2; k++ {
				v2[k] = m2[k][j]
			}
			m3[i][j], _ = v1.Dot(v2)
		}
	}
	return

}

func (m1 Matrix) Transpose() (m2 Matrix) {
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m2, _ = NewMartix(c, r)
	for i := 0; i < r; i++ {
		for j := 0; j < r; j++ {
			m2[i][j] = m1[i][j]
		}
	}
	return
}
