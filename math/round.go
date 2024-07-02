// Description: this module contains additional math functions
package math

import "math"

/*
round to a certain precision
*/
func RoundFloat(val float64, precision uint) float64 {
	// round to a certain precision
	ratio := math.Pow(10, float64(precision))
	// round to the nearest integer
	return math.Round(val*ratio) / ratio
}
