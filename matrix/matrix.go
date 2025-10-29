package matrix

// DiaganolOrder traverses a 2D matrix in diagonal order and returns elements as a slice.
// The function processes diagonals from bottom-left to top-right, where each diagonal
// is identified by the difference (column - row). Elements within each diagonal are
// collected and appended to the result in the order they appear.
// used recursive DFS or similar where visiting neighbors in all directions is valid
//
// Parameters:
//   - matrix: A 2D slice of integers with at least one row and one column
//
// Returns:
//   - A slice containing all matrix elements in diagonal order
//
// Example:
//
//	Input:  [[1,2,3], [4,5,6], [7,8,9]]
//	Output: [7,4,8,1,5,9,2,6,3]
func DiaganolOrder(matrix [][]int) []int {

	R := len(matrix)    //rows
	C := len(matrix[0]) //columns

	//diagnoal map stores elments
	diagMap := make(map[int][]int)

	// step 1 populate the map goig through the matrix
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ { //diagonal index based on difference
			D := c - r // diagonal index based on the difference
			diagMap[D] = append(diagMap[D], matrix[r][c])
			//append the current elment to the list for diaagional D
		}
	}

	//MinD occurs at (r=R-1, c=0) => D=0 - (R-1 = 1-R)
	//step 2 detemin the range of the digonal indicies
	minD := 1 - R

	//NaxD occurs at (r=R-0, c=C-1) => D = (C-1) - 0 = C-1
	maxD := C - 1

	//initialize the result slice
	res := make([]int, 0, R*C)

	// step 3 interage through diaganionalis in cincreasing order of D

	for D := minD; D <= maxD; D++ {

		if diagElements, ok := diagMap[D]; ok {
			res = append(res, diagElements...)
		}
	}
	return res
}
