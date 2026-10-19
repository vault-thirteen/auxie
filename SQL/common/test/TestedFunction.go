package ct

import "testing"

type TestedFunction struct {
	Name string
	Func func(t *testing.T)
}
