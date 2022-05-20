package square

import (
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {
	cases := map[string]struct {
		start        Point
		side         uint
		expEnd       Point
		expArea      uint
		expPerimeter uint
	}{
		"positive start": {
			start:        Point{3, 5},
			side:         4,
			expEnd:       Point{7, 9},
			expArea:      16,
			expPerimeter: 16,
		},
		"negative start": {
			start:        Point{-432, -678},
			side:         128,
			expEnd:       Point{-304, -550},
			expArea:      16384,
			expPerimeter: 512,
		},
		"zero side": {
			start:        Point{312, 534},
			side:         0,
			expEnd:       Point{312, 534},
			expArea:      0,
			expPerimeter: 0,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			sq := Square{
				start: tt.start,
				a:     tt.side,
			}

			end := sq.End()
			area := sq.Area()
			perimeter := sq.Perimeter()

			if !reflect.DeepEqual(end, tt.expEnd) {
				t.Errorf("the end Point is incorrect: exp: %v, got %v", tt.expEnd, end)
			}

			if area != tt.expArea {
				t.Errorf("the area is incorrect: exp: %d, got %d", tt.expArea, area)
			}

			if perimeter != tt.expPerimeter {
				t.Errorf("the perimeter is incorrect: exp: %d, got %d", tt.expPerimeter, perimeter)
			}

		})
	}
}
