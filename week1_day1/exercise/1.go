package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	row int
	col int
	mat [][]int
}

func (m *Matrix) getRows() int {
	return m.row
}

func (m *Matrix) getCols() int {
	return m.col
}

func (m *Matrix) setElement(i, j, x int) {
	m.mat[i][j] = x
}

func (m *Matrix) addMatrix(m1 Matrix) Matrix {
	res := *m
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			res.mat[i][j] += m1.mat[i][j]
		}
	}
	return res
}

func (m *Matrix) printAsJson() {
	matJson, _ := json.Marshal(m.mat)
	fmt.Println(string(matJson))
}

func main() {
	mat1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	mat2 := [][]int{{11, 21, 31}, {41, 51, 61}, {71, 81, 91}}
	m1 := Matrix{3, 3, mat1}
	m2 := Matrix{3, 3, mat2}

	fmt.Println("Rows:", m1.getRows())
	fmt.Println("Cols:", m1.getCols())
	fmt.Println("Initial:", m1.mat)
	m1.setElement(1, 1, -5)
	fmt.Println("Modified:", m1.mat)
	fmt.Println("Added matrix:", m1.addMatrix(m2))
	fmt.Println("As JSON:")
	m1.printAsJson()
}
