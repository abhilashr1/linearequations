package compute

//Determinant computes the value of the Determinant 2x2 matrix
func Determinant(mat [2][2]float32) float32 {
	var ans float32
	ans = mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	return ans
}

//Coeffmatrix builds the determinant matrix according to Cramers rule
func Coeffmatrix(mat [2][3]float32) ([2][2]float32, [2][2]float32, [2][2]float32) {
	var d, d1, d2 [2][2]float32

	d[0][0] = mat[0][0]
	d[0][1] = mat[0][1]
	d[1][0] = mat[1][0]
	d[1][1] = mat[1][1]

	d1[0][0] = mat[0][2]
	d1[0][1] = mat[0][1]
	d1[1][0] = mat[1][2]
	d1[1][1] = mat[1][1]

	d2[0][0] = mat[0][0]
	d2[0][1] = mat[0][2]
	d2[1][0] = mat[1][0]
	d2[1][1] = mat[1][2]

	return d, d1, d2
}

// Solution is the main compute function which uses Cramer's set of determinant matrices to find the solution
// in this case, of a system of linear equations with 2 variable.
// Cramer's rule is an explicit formula for the solution of a system of linear equations with as many equations as unknowns,
// valid whenever the system has a unique solution
func Solution(d [2][2]float32, d1 [2][2]float32, d2 [2][2]float32) (float32, float32) {
	det := Determinant(d)
	det1 := Determinant(d1)
	det2 := Determinant(d2)

	if det != 0 {
		val1 := det1 / det
		val2 := det2 / det
		return val1, val2
	}
	return -9999, -9999
}
