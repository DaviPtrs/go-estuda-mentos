package space

type Planet string

const EarthYearsSeconds int = 31557600

func getPlanetOrbital(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	case "Earth":
		return 1
	default:
		return -1
	}
}

func Age(seconds float64, planet Planet) float64 {
	orbitalValue := getPlanetOrbital(planet)
	if orbitalValue == -1 {
		return orbitalValue
	}
	return (seconds / float64(EarthYearsSeconds)) / orbitalValue
}
