package space

var orbitalPeriodsInEarthYears = map[string]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age returns the age on any planet given an age in seconds
func Age(seconds float64, planet string) float64 {
	return convertToEarthYears(seconds) / orbitalPeriodsInEarthYears[planet]
}

//convertToEarthYears turns seconds into years on earth
func convertToEarthYears(seconds float64) float64 {
	return seconds / 31557600.0
}
