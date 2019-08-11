package tax_test

import (
	"fmt"
	"testing"

	"github.com/weissm2040/meetup/unittest/tax"
)

func TestCalculator(t *testing.T) {
	type scenario struct {
		income      float64
		expectedTax float64
		expectedErr error
	}
	scenarios := []scenario{
		{-1, 0, fmt.Errorf("invalid income -1")},
		{0, 0, nil},
		{18200, 0, nil},
		{18201, 0.19, nil},
		{26000, 1482, nil},
		{37000, 3572, nil},
		{37001, 3572.325, nil},
	}

	c := tax.NewCalculator()

	for _, s := range scenarios {
		tax, err := c.Calculate(s.income)
		if s.expectedErr != nil && err.Error() != s.expectedErr.Error() {
			t.Errorf("For '%v' income, expected error: '%v' but got: '%v'", s.income, s.expectedErr, err)
		}
		if tax != s.expectedTax {
			t.Errorf("For '%v' income, expected tax: '%v' but got: '%v'", s.income, s.expectedTax, tax)
		}
	}
}
