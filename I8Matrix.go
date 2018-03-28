package la4go

type I8Matrix struct {
	n    int
	m    int
	data [][]int8
}

func NewI8Matrix(n, m int) I8Matrix {
	data := make([][]int8, n)

	for i := 0; i < n; i++ {
		data[i] = make([]int8, m)
	}

	return I8Matrix{n, m, data}
}
