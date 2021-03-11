package nurbs

import "fmt"

/* findSpan : determines the knot span index */
/* knot span :The range of parameter values between two successive knots in a spline. */
/* 	input:
--n : last index in control points vector
--p : degree
--u : the variable that lies on the knot span
--U : the knot vector
*/
func findSpan(n, p int, u float64, U []int) int {
	if u == float64(U[n+1]) {
		return n
	}
	/* do binary search */
	low := p
	high := n + 1
	mid := (low + high) / 2
	/* BUG: this is not robust: */
	for u < float64(U[mid]) || u >= float64(U[mid+1]) {
		fmt.Println("knot Vector:", U)

		if u < float64(U[mid]) {
			high = mid
		} else {
			low = mid
		}
		fmt.Println("mid:", mid)
		mid = (low + high) / 2
	}
	return mid
}

/* basisFuns : computes the non vanishing basis functions */
/* in a dynamic programming fashion , this algorithm stores the
---results we computed by introducing the left and right notations */
/* input :
--i : the span index
--u : the variable that lies on the knot span
--p : the degree of the polynomial
--U : the knot vector
*/
func basisFuns(i int, u float64, p int, U []int) []float64 {
	N := make([]float64, p+1)
	left := make([]float64, p+1)
	right := make([]float64, p+1)

	N[0] = 1.
	for j := 1; j <= p; j++ {
		left[j] = u - float64(U[i+1-j])
		right[j] = float64(U[i+j]) - u
		saved := 0.
		for r := 0; r < j; r++ {
			temp := N[r] / (right[r+1] + left[j-r])
			N[r] = saved + right[r+1]*temp
			saved = left[j-r] * temp
		}
		N[j] = saved
	}
	return N
}

/* dersBasisFuns : computes nonzero basis functoins and their
---derivates. first step is the `basisFuns` modified tp store
---functions and and knot differences
*/
/* input:
-- i : the span index
-- u : the variable that lies on the knot span
-- p : the degree of the polynomial
-- n :
-- U : the knot vector
*/
/* output:
-- ders : the derivatives of the function
*/
func dersBasisFuns(i, u, p, n int, U []int) [][]float64 {
	/* ndu stores the basis functions and knot differences */
	ndu := make([][]float64, p+1)
	ders := make([][]float64, p+1)
	a := make([][]float64, p+1)

	for entry := range ndu {
		ndu[entry] = make([]float64, p+1)
		ders[entry] = make([]float64, p+1)
		a[entry] = make([]float64, p+1)
	}

	left := make([]float64, p+1)
	right := make([]float64, p+1)
	/* ------------------------------------ */
	for j := 1; j <= p; j++ {
		left[j] = float64(u - U[i+1-j])
		right[j] = float64(U[i+1] - u)
		saved := 0.
		for r := 0; r < j; r++ {
			/* lower triangle */
			ndu[j][r] = right[r+1] + left[j-r]
			temp := ndu[r][j-1] / ndu[j][r]
			/* upper triangle */
			ndu[r][j] = saved + right[r+1]*temp
			saved = left[j-r] * temp
		}
		ndu[j][j] = saved
	}
	/* ------------------------------------ */
	/* loading the basis functions */
	for j := 0; j <= p; j++ {
		ders[0][j] = ndu[j][p]
	}
	/* this section computes the derivatives: */
	/* ------------------------------------ */
	/* loop over function index */
	for r := 0; r <= p; r++ {
		/* define alternative rows in array a */
		s1 := 0
		s2 := 1
		a[0][0] = 1.
		/* loop to compute kth derivative */
		for k := 1; k <= n; k++ {
			d := 0.
			rk := r - k
			pk := p - k
			if r >= k {
				a[s2][0] = a[s1][0] / ndu[pk+1][rk]
				d = a[s2][0] * ndu[rk][pk]
			}
			j1 := 0
			j2 := 0
			if rk >= -1 {
				j1 = 1
			} else {
				j1 = -rk
			}
			if r-1 <= pk {
				j2 = k - 1
			} else {
				j2 = p - r
			}
			for j := j1; j <= j2; j++ {
				a[s2][j] = (a[s1][j] - a[s1][j-1]) / ndu[pk+1][rk+j]
				d += a[s2][j] * ndu[r][pk]
			}
			if r <= pk {
				a[s2][k] = -a[s1][k-1] / ndu[pk+1][r]
				d += a[s2][k] * ndu[r][pk]
			}
			ders[k][r] = d
			/* switching rows */
			j := s1
			s1 = s2
			s2 = j
		}
	}
	/* multiply through by the correct factors */
	r := p
	for k := 1; k <= n; k++ {
		for j := 0; j <= p; j++ {
			ders[k][j] *= float64(r)
		}
		r *= (p - k)
	}
	return ders
}
