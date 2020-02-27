package main

func LevenshteinDistance(x, y string) int {
	if len(x) == 0 {
		return len(y)
	} else if len(y) == 0 {
		return len(x)
	}

	matrix := make([][]int, len(x)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(y)+1)
	}

	for i := range matrix[0] {
		matrix[0][i] = i
	}

	for i := range matrix {
		matrix[i][0] = i
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {

			var min = matrix[i-1][j-1]

			if x[i-1] != y[j-1] {
				min++
			}

			if matrix[i-1][j]+1 < min {
				min = matrix[i-1][j] + 1
			} else if matrix[i][j-1]+1 < min {
				min = matrix[i][j-1] + 1
			}

			matrix[i][j] = min

		}
	}
	return matrix[len(matrix)-1][len(matrix[0])-1]
}
