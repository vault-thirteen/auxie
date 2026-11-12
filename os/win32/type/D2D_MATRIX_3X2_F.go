package t

import (
	b "github.com/vault-thirteen/auxie/os/win32/base"
)

type D2D_MATRIX_3X2_F struct {
	// Horizontal scaling / cosine of rotation
	M11 b.FLOAT

	// Vertical shear / sine of rotation
	M12 b.FLOAT

	// Horizontal shear / negative sine of rotation
	M21 b.FLOAT

	// Vertical scaling / cosine of rotation
	M22 b.FLOAT

	// Horizontal shift (always orthogonal regardless of rotation)
	Dx b.FLOAT

	// Vertical shift (always orthogonal regardless of rotation)
	Dy b.FLOAT
}

// Array returns the matrix data as a 3x2 array,
// mimicking the C-style FLOAT m[3][2] union field.
func (m *D2D_MATRIX_3X2_F) Array() [3][2]b.FLOAT {
	return [3][2]b.FLOAT{
		{m.M11, m.M12},
		{m.M21, m.M22},
		{m.Dx, m.Dy},
	}
}

// FromArray populates the matrix from a 3x2 array.
func (m *D2D_MATRIX_3X2_F) FromArray(a [3][2]b.FLOAT) {
	m.M11, m.M12 = a[0][0], a[0][1]
	m.M21, m.M22 = a[1][0], a[1][1]
	m.Dx, m.Dy = a[2][0], a[2][1]
}

func D2D_Matrix3x2F_Identity() D2D_MATRIX_3X2_F {
	return D2D_MATRIX_3X2_F{
		M11: 1, M12: 0,
		M21: 0, M22: 1,
		Dx: 0, Dy: 0,
	}
}
