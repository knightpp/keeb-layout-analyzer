package distance

import "fmt"

type Distance int

const (
	MM Distance = 1
	CM          = 10 * MM
	M           = 100 * CM
)

func (d Distance) String() string {
	if d < CM {
		return fmt.Sprintf("%dmm", d)
	}
	if d < M {
		return fmt.Sprintf("%dcm %dmm", d/CM, d%CM)
	}
	return fmt.Sprintf("%dm %dcm %dmm", d/M, d%M/CM, d%M%CM)
}
