package functionaloption

import (
	"math"
	"testing"
)

func TestDo(t *testing.T) {
	var c Calculator
	var result float64

	resultErrFmt := func(op string, v float64) {
		t.Fatalf("expected result of %s to equal %f, got %f", op, v, result)
	}

	result = c.Do(Add(10))
	if result != 10 {
		resultErrFmt("add", 10)
	}

	result = c.Do(Sub(8))
	if result != 2 {
		resultErrFmt("sub", 2)
	}

	result = c.Do(Mul(2))
	if result != 4 {
		resultErrFmt("mul", 4)
	}

	result = c.Do(math.Sqrt)
	if result != 2 {
		resultErrFmt("sqrt", 2)
	}
}
