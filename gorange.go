package gorange

import (
	"fmt"
	"math"
)

type Gorange struct {
	idealVal     float64
	radius       float64
	displayWidth int
}

// New creates a new Gorange instance, where idealVal will be the perfect value
// in the middle of the range, and radius is the size of the inner (valid) range.
func New(idealVal float64, radius float64) *Gorange {
	return &Gorange{
		idealVal:     idealVal,
		radius:       radius,
		displayWidth: 21,
	}
}

// SetDisplayWidth lets you modify the expected string width.
// Valid values are: 5, 9, 13, ...
//
// Examples:
//
//	5: -[|]-
//	9: --[-|-]--
//	13: ---[--|--]---
//
// Default width is 21.
func (r *Gorange) SetDisplayWidth(width int) error {
	if width < 5 || width%4 != 1 {
		return fmt.Errorf("invalid width value: %d, correct values: 5, 9, 13 etc", width)
	}

	r.displayWidth = width
	return nil
}

// GetRangeAsText return a string of range with given value
// marked as "█" if it's within or close enough to the range.
//
// Example: "-----[----█----]-----"
func (r *Gorange) GetRangeAsText(value float64) string {
	str := r.getEmptyRange()
	markIndex := r.getMarkIndex(value)

	if markIndex > -1 {
		str = str[:markIndex] + "█" + str[markIndex+1:]
	}

	return str
}

func (r *Gorange) getEmptyRange() string {
	str := "-"
	extraDashesWidth := (r.displayWidth - 5) / 4
	for i := 0; i < extraDashesWidth; i++ {
		str += "-"
	}
	str += "["
	for i := 0; i < extraDashesWidth; i++ {
		str += "-"
	}
	str += "|"
	for i := 0; i < extraDashesWidth; i++ {
		str += "-"
	}
	str += "]"
	for i := 0; i < extraDashesWidth; i++ {
		str += "-"
	}
	str += "-"

	return str
}

func (r *Gorange) getMarkIndex(value float64) int {
	stepDiff := r.radius / (float64(r.displayWidth-5)/4 + 1)
	vals := make([]float64, r.displayWidth)
	for i := 0; i < len(vals); i++ {
		vals[i] = r.idealVal - stepDiff*float64((r.displayWidth-1)/2-i)
	}

	if value < vals[0]-(stepDiff/2) || value >= vals[len(vals)-1]+(stepDiff/2) {
		// out of range
		return -1
	}

	return findClosestValueIndex(vals, value)
}

func findClosestValueIndex(vals []float64, val float64) int {
	closestIndex := 0
	closestDiff := math.Abs(val - vals[0])

	for i := 1; i < len(vals); i++ {
		diff := math.Abs(val - vals[i])
		if diff < closestDiff || (diff == closestDiff && vals[i] > vals[closestIndex]) {
			closestDiff = diff
			closestIndex = i
		}
	}

	return closestIndex
}
