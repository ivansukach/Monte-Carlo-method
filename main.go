package main

import (
	"github.com/ivansukach/Monte-Carlo-method/generators"
	"github.com/ivansukach/Monte-Carlo-method/statistics"
	"github.com/sirupsen/logrus"
	"math"
)

func ksi1(x float64) float64 {
	return math.Pi * math.Cos(x) * math.Cos(x) * x * x
}
func ksi2(x float64) float64 {
	return math.Exp(math.Pow(x, 2)/2) * math.Sqrt(2*math.Pi) / (x*x + x + 1)
}

func main() {
	a1 := 0.
	b1 := math.Pi
	a01 := 296454621
	a02 := 302711857
	c1 := 48840859
	c2 := 37330745
	M := int(math.Pow(2, 31))
	K := 64
	n := 1000
	logrus.Info("M: ", M)
	aSequence2 := *generators.LinearCongruential(a02, c2, M, n)
	aSequence1SpecialSize := *generators.LinearCongruential(a01, c1, M, n+K)
	sequenceByMacLarenMarsaglia := *generators.MacLarenMarsaglia(aSequence1SpecialSize, aSequence2, K, n)
	logrus.Info("Value of integral1 : a[1]=",
		statistics.E(ksi1, n, *statistics.UniformDistributionVariates(a1, b1, sequenceByMacLarenMarsaglia)))
	logrus.Info("Value of integral2 : a[2]=",
		statistics.E(ksi2, n, *statistics.NormalDistributionVariates(sequenceByMacLarenMarsaglia)))
	logrus.Info("Value of integral2 : a[2]=",
		statistics.E(ksi2, n, *statistics.AlternativeNormalVariates(sequenceByMacLarenMarsaglia)))
}
