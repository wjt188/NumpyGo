package matrix

type RowVector []float64
type ColumVector [][1]float64
type Matrix [][]float64

type IVecMat interface {
	GetShape() (s [2]int)
	Mul(c float64)
}

type IRowVec interface {
	Add(rv2 RowVector) (rv3 RowVector, err error)   //行向量相加
	Minus(rv2 RowVector) (rv3 RowVector, err error) //行向量相减
	Dot(rv2 RowVector) (dot float64, err error)     //行向量点乘
	Cross(rv2 RowVector) (rv3 RowVector, err error)
	Length() (l float64)
	Transpose() (cv ColumVector) //矩阵转置
}

type IMat interface {
	Add(m2 Matrix) (m3 Matrix, err error)
	Minus(m2 Matrix) (m3 Matrix, err error)
	MatMul(m2 Matrix) (m3 Matrix, err error)
	Transpose() (m2 Matrix)
}
