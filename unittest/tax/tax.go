package tax

import "fmt"

type taxBand struct {
	upper float64
	rate  float64
}

// Calculator can calculate tax for a given income.
type Calculator struct {
	taxBands []taxBand
}

// NewCalculator creates a new tax calculator loaded with the current tax bands.
func NewCalculator() Calculator {
	return Calculator{
		taxBands: []taxBand{
			{18200, 0.0},
			{37000, 0.19},
			{90000, 0.325},
			{180000, 0.37},
			{999999999, 0.45},
		},
	}
}

// Calculate will calculate the amount of tax to be paid given an income.
func (t Calculator) Calculate(income float64) (float64, error) {
	if income < 0 {
		return 0, fmt.Errorf("invalid income %v", income)
	}
	tax := 0.0
	processed := 0.0
	for _, band := range t.taxBands {
		process := min(income, band.upper)
		tax += (process - processed) * band.rate
		if income <= band.upper {
			break
		}
		processed = band.upper

	}
	return tax, nil
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
