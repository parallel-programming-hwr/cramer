package main

import (
	"fmt"
	"os"
)

func main() {

	matrix := [][]int{{2, 3, 5, 2}, {1, 3, 2, 5}, {2, 5, 1, 6}, {1, 5, 2, 7}}
	b := []int{2, 2, 7, 8}
	checkMatAndVec(matrix, b)
	fmt.Printf("Matrix: \n")
	outMat(matrix)
	checkMat(matrix)
	detM := getDet(matrix)
	x := make([]float32, len(matrix))
	var ch = make(chan int)
	for i := 0; i < len(matrix); i++ {
		go calcX(matrix, x, b, i, detM, ch)
	}
	fmt.Printf("\nLoesungsvektor: \n")
	for l := 0; l < len(x); l++ {
		<-ch
	}
	for k := 0; k < len(x); k++ {
		fmt.Printf("x%d = %f \n", k, x[k])
	}
}

func calcX(matrix [][]int, x []float32, b []int, i int, detM int, ch chan int) {
	matai := copyMat(matrix)
	for j := 0; j < len(matrix[i]); j++ {
		matai[j][i] = b[j]
	}
	detai := getDet(matai)
	//fmt.Printf("\nMatrix %d\n", i)
	//outMat(matai)
	//fmt.Printf("Determinante %d => x%d = %f\n", detai, i, float32(getDet(matai))/float32(detM))
	x[i] = float32(detai) / float32(detM)
	ch <- 1
}

func checkMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		if len(mat) != len(mat[i]) {
			fmt.Printf("Matrix ungueltig")
			os.Exit(1)
		}

	}
}

func checkMatAndVec(matrix [][]int, b []int) {
	checkMat(matrix)
	if len(matrix) != len(b) {
		fmt.Printf("Matrix und vektor nicht gleich groß")
		os.Exit(1)
	}
}

func getDet(mat [][]int) int {
	if len(mat) > 1 {
		erg := 0
		neg := false
		var ch = make(chan int)
		for i := 0; i < len(mat); i++ {
			go singleDet(mat, neg, i, ch)
			neg = !neg
		}
		for i := 0; i < len(mat); i++ {
			zw := <-ch
			erg = erg + zw
		}
		return erg
	} else {
		return mat[0][0]
	}

}

func singleDet(mat [][]int, neg bool, i int, ch chan int) {
	erg := 0
	mat1 := [][]int{}
	for j := 0; j < len(mat); j++ {
		if j == i {
			continue
		}
		mat1 = append(mat1, []int{})
		for k := 1; k < len(mat[j]); k++ {
			if j < i {
				mat1[j] = append(mat1[j], mat[j][k])
			} else {
				mat1[j-1] = append(mat1[j-1], mat[j][k])
			}
		}

	}
	if !neg {
		erg = erg + mat[i][0]*getDet(mat1)
	} else {
		erg = erg - mat[i][0]*getDet(mat1)
	}
	ch <- erg
}

func outMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Printf("\n")
	}
}

func copyMat(mat [][]int) [][]int {
	matai := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		matai[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			matai[i][j] = mat[i][j]
		}
	}
	return matai
}
