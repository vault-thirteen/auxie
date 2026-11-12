package t

import (
	"testing"

	b "github.com/vault-thirteen/auxie/os/win32/base"
	"github.com/vault-thirteen/auxie/tester"
)

func Test_D2D_MATRIX_3X2_F_Array(t *testing.T) {
	aTest := tester.New(t)

	m := D2D_MATRIX_3X2_F{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	expectedResult := [3][2]b.FLOAT{
		{1.1, 2.2},
		{3.3, 4.4},
		{5.5, 6.6},
	}
	aTest.MustBeEqual(m.Array(), expectedResult)
}

func Test_D2D_MATRIX_3X2_F_FromArray(t *testing.T) {
	aTest := tester.New(t)

	x := [3][2]b.FLOAT{
		{1.1, 2.2},
		{3.3, 4.4},
		{5.5, 6.6},
	}
	var m D2D_MATRIX_3X2_F
	m.FromArray(x)
	expectedResult := D2D_MATRIX_3X2_F{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	aTest.MustBeEqual(m, expectedResult)
}

func Test_D2D_MATRIX_3X2_F_Identity(t *testing.T) {
	aTest := tester.New(t)

	expectedResult := D2D_MATRIX_3X2_F{
		M11: 1, M12: 0,
		M21: 0, M22: 1,
		Dx: 0, Dy: 0,
	}
	aTest.MustBeEqual(D2D_Matrix3x2F_Identity(), expectedResult)
}
