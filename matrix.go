package main

type m64 [][]float64

func newM64(rows, cols int) m64 {

	result := make(m64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	return result
}

func (m m64) Add(m2 m64) m64 {

	row := len(m)
	for i := 0; i < row; i++ {
		col := len(m[i])
		for j := 0; j < col; j++ {
			m[i][j] = m[i][j] + m2[i][j]
		}
	}

	return m
}

func (m m64) Sub(m2 m64) m64 {

	row := len(m)
	for i := 0; i < row; i++ {
		col := len(m[i])
		for j := 0; j < col; j++ {
			m[i][j] = m[i][j] - m2[i][j]
		}
	}

	return m
}

func (m m64) MMul(m2 m64) m64 {

	m1Rows := len(m)
	//m1Cols := len(m[0])
	m2Rows := len(m2)
	m2Cols := len(m2[0])

	result := newM64(m1Rows, m2Cols)

	for i := 0; i < m1Rows; i++ {

		for j := 0; j < m2Cols; j++ {

			var sum float64
			for k := 0; k < m2Rows; k++ {
				sum += m[i][k] * m2[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}

func (m m64) SMul(m2 m64) m64 {
	row := len(m)
	for i := 0; i < row; i++ {
		col := len(m[i])
		for j := 0; j < col; j++ {
			m[i][j] = m[i][j] * m2[i][j]
		}
	}

	return m
}

func (m m64) Transpose() m64 {
	rows := len(m)
	cols := len(m[0])

	result := newM64(cols, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = m[i][j]
		}
	}

	return result
}

func (m m64) ApplyFunc(f func(float64) float64) m64 {
	row := len(m)
	for i := 0; i < row; i++ {
		col := len(m[i])
		for j := 0; j < col; j++ {
			m[i][j] = f(m[i][j])
		}
	}

	return m
}
