package statistics

import "math"

func E(ksi func(x float64) float64, n int, a []float64) float64 {
	sum := 0.
	for i := 0; i < n; i++ {
		sum += ksi(a[i])
	}
	return sum / float64(n)
}
func UniformDistributionVariates(a float64, b float64, basicVariate []float64) *[]float64 {
	n := len(basicVariate)
	uniformVariates := make([]float64, n)
	for i := 0; i < n; i++ {
		uniformVariates[i] = a + (b-a)*basicVariate[i]
	}
	return &uniformVariates
}
func NormalDistributionVariates(basicVariate []float64) *[]float64 {
	n := len(basicVariate)
	normalVariates := make([]float64, n)
	for i := 0; i < n-1; i++ {
		normalVariates[i] = 2 * math.Sqrt(-2*math.Log(basicVariate[i])) * math.Cos(2*math.Pi*basicVariate[i+1])
	}
	normalVariates[n-1] = 2 * math.Sqrt(-2*math.Log(basicVariate[n-2])) * math.Sin(2*math.Pi*basicVariate[n-1])
	return &normalVariates
}
func AlternativeNormalVariates(basicVariate []float64) *[]float64 {
	n := len(basicVariate)
	normalVariates := make([]float64, n)
	sqrt2pi := math.Sqrt(2 * math.Pi)
	for i := 0; i < n; i++ {
		tmp := 2 * math.Log(sqrt2pi*basicVariate[i])
		if tmp < 0 {
			normalVariates[i] = -math.Sqrt(-tmp)
		} else {
			normalVariates[i] = math.Sqrt(tmp)
		}
	}
	return &normalVariates
}
