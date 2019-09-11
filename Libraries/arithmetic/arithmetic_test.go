package arithmetic

import "testing"

func TestSum(t *testing.T) {
	a := 10
	b := 20
	e := 30
	r := Sum(a, b)

	if r != e {
		t.Errorf("Incorrect sum result, %d + %d is not %d", a, b, r)
	}

}
