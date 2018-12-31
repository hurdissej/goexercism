package space

import (
	"math"
)

//Age returns the age on any planet given an age in seconds
func Age(seconds float64, planet string) float64 {
	switch planet {
	case "Earth":
		return ConvertToEarthYears(seconds)
	case "Mercury":
		return 1
	case "Venus":
		return 1
	case "Mars":
		return 1
	case "Jupiter":
		return 1
	case "Saturn":
		return 1
	case "Uranus":
		return 1
	case "Neptune":
		return 1
	}
	return -1
}

//ConvertToEarthYears turns seconds into years on earth
func ConvertToEarthYears(seconds float64) float64 {
	return math.Round(((seconds / (365 * 24 * 60 * 60)) * 10.05) / 10)
}
