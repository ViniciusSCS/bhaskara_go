package bhaskara

import "math"

func Bhaskara(a, b, c float64) (float64, float64, bool) {
	delta := b*b - 4*a*c

	if delta < 0 {
		return 0, 0, false // Não há raízes reais
	}

	raizDelta := math.Sqrt(delta)
	x1 := (-b + raizDelta) / (2 * a)
	x2 := (-b - raizDelta) / (2 * a)

	return x1, x2, true
}
