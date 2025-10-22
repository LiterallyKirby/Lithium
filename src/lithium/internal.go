package lithium

import "math"

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func sine(a float64) float64 {
	return math.Sin(a)
}

func cosine(a float64) float64 {
	return math.Cos(a)
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}
