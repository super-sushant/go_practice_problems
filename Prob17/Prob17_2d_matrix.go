package main

import (
	"fmt"
)

// create matrix where ele = i*j
func createMatrix(rows,cols int) ([][]int) {
	
	mat:= make([][]int,rows)

	for i:=0; i<rows;i++ {
		// make empty slice
		row:=make([]int, cols)

		for j:=0; j<cols; j++ {
			// append in slice
			row[j]= i*j
		}

		mat[i]= row
	}

	return mat
	
}

func main(){

	mat:=createMatrix(7,7)

	fmt.Println(mat)
}