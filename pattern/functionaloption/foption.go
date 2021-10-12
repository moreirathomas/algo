package functionaloption

// Calculator represents a calculator object. It has an internal
// `acc` field used as a destination for the operations it handles.
type Calculator struct {
	acc float64
}

// Do executes the given operation and stores the result inside the calculator.
func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func Add(n float64) func(float64) float64 {
	return func(a float64) float64 { return a + n }
}
func Sub(n float64) func(float64) float64 {
	return func(a float64) float64 { return a - n }
}
func Mul(n float64) func(float64) float64 {
	return func(a float64) float64 { return a * n }
}

// Div() would require to change the signature as dividing by 0 should be an error.
