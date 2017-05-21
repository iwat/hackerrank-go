package dfs

type Matrix struct {
	matrix []byte
	rows   int
	cols   int
}

func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{make([]byte, rows*cols), rows, cols}
}

func (m *Matrix) FillRow(row int, value ...byte) {
	offset := row * m.cols
	for i := 0; i < m.cols; i++ {
		m.matrix[offset+i] = value[i]
	}
}

func (m *Matrix) LargestRegionSize() int {
	visited := make([]bool, len(m.matrix))
	largest := 0
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			size := dfs(m, row, col, visited)
			if size > largest {
				largest = size
			}
		}
	}
	return largest
}

func dfs(m *Matrix, row, col int, visited []bool) int {
	if row < 0 || col < 0 || row >= m.rows || col >= m.cols {
		return 0
	}
	offset := row*m.cols + col
	if visited[offset] {
		return 0
	}
	visited[offset] = true
	if m.matrix[offset] == 0 {
		return 0
	}
	size := 1
	size += dfs(m, row-1, col-1, visited)
	size += dfs(m, row-1, col, visited)
	size += dfs(m, row-1, col+1, visited)
	size += dfs(m, row, col-1, visited)
	size += dfs(m, row, col+1, visited)
	size += dfs(m, row+1, col-1, visited)
	size += dfs(m, row+1, col, visited)
	size += dfs(m, row+1, col+1, visited)
	return size
}
