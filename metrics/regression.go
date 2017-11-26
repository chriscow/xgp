package metrics

import (
	"math"
)

// MeanAbsoluteError measures the mean absolute error (MAE).
type MeanAbsoluteError struct{}

// Apply MeanAbsoluteError.
func (metric MeanAbsoluteError) Apply(yTrue, yPred, weights []float64) (float64, error) {
	if len(yTrue) != len(yPred) {
		return math.Inf(1), &errMismatchedLengths{len(yTrue), len(yPred)}
	}
	if weights != nil && len(yTrue) != len(weights) {
		return math.Inf(1), &errMismatchedLengths{len(yTrue), len(weights)}
	}

	var (
		sum float64
		ws  float64
	)
	for i := range yTrue {
		if weights != nil {
			sum += math.Abs(yTrue[i]-yPred[i]) * weights[i]
			ws += weights[i]
		} else {
			sum += math.Abs(yTrue[i] - yPred[i])
		}
	}
	if weights != nil {
		return sum / ws, nil
	}
	return sum / float64(len(yTrue)), nil
}

// Classification method of MeanAbsoluteError.
func (metric MeanAbsoluteError) Classification() bool {
	return false
}

// BiggerIsBetter method of MeanAbsoluteError.
func (metric MeanAbsoluteError) BiggerIsBetter() bool {
	return false
}

// NeedsProbabilities method of MeanAbsoluteError.
func (metric MeanAbsoluteError) NeedsProbabilities() bool {
	return false
}

// String method of MeanAbsoluteError.
func (metric MeanAbsoluteError) String() string {
	return "mae"
}

// MeanSquaredError measures the mean absolute error (MAE).
type MeanSquaredError struct{}

// Apply MeanSquaredError.
func (metric MeanSquaredError) Apply(yTrue, yPred, weights []float64) (float64, error) {
	if len(yTrue) != len(yPred) {
		return math.Inf(1), &errMismatchedLengths{len(yTrue), len(yPred)}
	}
	if weights != nil && len(yTrue) != len(weights) {
		return math.Inf(1), &errMismatchedLengths{len(yTrue), len(weights)}
	}

	var (
		sum float64
		ws  float64
	)
	for i := range yTrue {
		if weights != nil {
			sum += math.Pow(yTrue[i]-yPred[i], 2) * weights[i]
			ws += weights[i]
		} else {
			sum += math.Pow(yTrue[i]-yPred[i], 2)
		}
	}
	if weights != nil {
		return sum / ws, nil
	}
	return sum / float64(len(yTrue)), nil
}

// Classification method of MeanSquaredError.
func (metric MeanSquaredError) Classification() bool {
	return false
}

// BiggerIsBetter method of MeanSquaredError.
func (metric MeanSquaredError) BiggerIsBetter() bool {
	return false
}

// NeedsProbabilities method of MeanSquaredError.
func (metric MeanSquaredError) NeedsProbabilities() bool {
	return false
}

// String method of MeanSquaredError.
func (metric MeanSquaredError) String() string {
	return "mse"
}

// R2 measures the coefficient of determination.
type R2 struct{}

// Apply R2.
func (metric R2) Apply(yTrue, yPred, weights []float64) (float64, error) {
	if len(yTrue) != len(yPred) {
		return math.Inf(1), &errMismatchedLengths{len(yTrue), len(yPred)}
	}
	if weights != nil && len(yTrue) != len(weights) {
		return math.Inf(1), &errMismatchedLengths{len(yTrue), len(weights)}
	}

	// Compute the mean of the observed data
	var (
		yMean float64
		ws    float64
	)
	for i, y := range yTrue {
		if weights != nil {
			yMean += y * weights[i]
			ws += weights[i]
		} else {
			yMean += y
		}
	}
	if weights != nil {
		yMean /= ws
	} else {
		yMean /= float64(len(yTrue))
	}

	var (
		SSR float64
		SST float64
	)
	for i := range yTrue {
		if weights != nil {
			SSR += math.Pow(yPred[i]-yTrue[i], 2) * weights[i]
			SST += math.Pow(yTrue[i]-yMean, 2) * weights[i]
		} else {
			SSR += math.Pow(yPred[i]-yTrue[i], 2)
			SST += math.Pow(yTrue[i]-yMean, 2)
		}
	}

	return 1 - SSR/SST, nil
}

// Classification method of R2.
func (metric R2) Classification() bool {
	return false
}

// BiggerIsBetter method of R2.
func (metric R2) BiggerIsBetter() bool {
	return false
}

// NeedsProbabilities method of R2.
func (metric R2) NeedsProbabilities() bool {
	return false
}

// String method of R2.
func (metric R2) String() string {
	return "r2"
}
