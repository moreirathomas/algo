package algo

type Calculator struct {
	acc float64
}

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
