package distance_test

import (
	"testing"

	"github.com/knightpp/keeb-layout-analyzer/internal/distance"
)

func TestString(t *testing.T) {
	tt := []struct {
		expect   string
		distance distance.Distance
	}{
		{
			expect:   "1mm",
			distance: distance.MM,
		},
		{
			expect:   "9mm",
			distance: 9 * distance.MM,
		},
		{
			expect:   "1cm 9mm",
			distance: 1*distance.CM + 9*distance.MM,
		},
		{
			expect:   "40m 1cm 9mm",
			distance: 40*distance.M + distance.CM + 9*distance.MM,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.expect, func(t *testing.T) {
			if tc.distance.String() != tc.expect {
				t.Fatalf("%s != %s", tc.distance, tc.expect)
			}
		})
	}
}
