package main

import (
	"fmt"

	"github.com/iwat/algorithms/dfs"
)

func main() {
	var rows int
	var cols int
	fmt.Scanf("%d", &rows)
	fmt.Scanf("%d", &cols)
	matrix := dfs.NewMatrix(rows, cols)
	for row := 0; row < rows; row++ {
		rowData := make([]byte, cols)
		for col := 0; col < cols; col++ {
			var data byte
			fmt.Scanf("%d", &data)
			rowData[col] = data
		}
		matrix.FillRow(row, rowData...)
	}
	fmt.Println(matrix.LargestRegionSize())
}
