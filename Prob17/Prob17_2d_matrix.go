package main

import (
	"fmt"
)

// create matrix where ele = i*j
func createMatrix(rows,cols int) ([][]int) {
	
	mat:= make([][]int,0)

	for i:=0; i<rows;i++ {
		// make empty slice
		row:=make([]int, 0)

		for j:=0; j<cols; j++ {
			// append in slice
			row=append(row, i*j)
		}
		
		mat=append(mat, row)
	}

	return mat
	
}

func main(){

	mat:=createMatrix(7,7)

	fmt.Println(mat)
}