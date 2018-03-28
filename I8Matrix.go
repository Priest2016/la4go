package la4go

import (
	"errors"
)

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

func (i8m *I8Matrix) GetN() int {
	return i8m.n
}

func (i8m *I8Matrix) GetM() int {
	return i8m.m
}

func (i8m *I8Matrix) GetData() [][]int8 {
	return i8m.data
}

func (i8m *I8Matrix) GetCell(n, m int) (int8, error) {
	if n <= i8m.n && m <= i8m.m {
		return i8m.data[n][m], nil
	}

	return int8(-1), errors.New("Cantfcell")
}

func (i8m *I8Matrix) GetRow(rowNum) ([]int8, error) {
	if rowNum <= i8m.n {
		return i8m.data[rowNum], nil
	}

	return make([]int8, 0), errors.New("Cantfrow")
}

func (i8m *I8Matrix) GetColumn(colNum int) ([]int8, error) {
	if colNum < i8m.m {
		rv := make([]int8, n)
		for i := 0; i < i8m.n; i++ {
			rv[i] = i8m.data[i][colNum]
		}

		return rv, nil
	}

	return make([]int8, 0), errors.New("Cantfcol")
}
