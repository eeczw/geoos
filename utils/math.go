package utils

func Clamp(x, lower, upper float64) float64 {
	if x > upper {
		return upper
	}
	if x < lower {
		return lower
	}
	return x
}
