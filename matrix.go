package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func diagonalDifference(matrix [][]int) int {
	n := len(matrix)
	primaryDiagonalSum := 0
	secondaryDiagonalSum := 0

	for i := 0; i < n; i++ {
		primaryDiagonalSum += matrix[i][i]
		secondaryDiagonalSum += matrix[i][n-i-1]
	}

	return int(math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum)))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan ukuran matriks (N x N):")
	sizeInput, _ := reader.ReadString('\n')
	sizeInput = strings.TrimSpace(sizeInput)
	n, err := strconv.Atoi(sizeInput)
	if err != nil {
		fmt.Println("Ukuran tidak valid:", err)
		return
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen-elemen baris ke-%d yang dipisahkan dengan spasi:\n", i+1)
		rowInput, _ := reader.ReadString('\n')
		rowInput = strings.TrimSpace(rowInput)
		rowStrArr := strings.Split(rowInput, " ")

		if len(rowStrArr) != n {
			fmt.Println("Jumlah elemen tidak sesuai dengan ukuran matriks")
			return
		}

		row := make([]int, n)
		for j, str := range rowStrArr {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Input tidak valid:", err)
				return
			}
			row[j] = num
		}
		matrix[i] = row
	}
	fmt.Println("Matriks integer:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	result := diagonalDifference(matrix)
	fmt.Println("Hasil pengurangan matrix diagonal adalah = ", result)
}
