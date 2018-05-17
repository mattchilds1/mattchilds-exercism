package space


type Planet string

func Age (s float64, planet Planet) float64 {
	var age float64
	var earthyearseconds float64 = 31557600
	var mercuryyearseconds float64 = earthyearseconds * 0.2408467
	var venusyearseconds float64 = earthyearseconds * 0.61519726
	var marsyearseconds float64 = earthyearseconds * 1.8808158
	var jupiteryearsseconds float64 = earthyearseconds * 11.862615
	var saturnyearseconds float64 = earthyearseconds * 29.447498
	var uranusyearseconds float64 = earthyearseconds * 84.016846
	var neptuneyearseconds float64 = earthyearseconds * 164.79132

	switch {
	case planet == "Earth" :
		age = s / earthyearseconds
	case planet == "Mercury" :
		age = s / mercuryyearseconds
	case planet == "Venus" :
		age = s / venusyearseconds
	case planet == "Mars" :
		age = s / marsyearseconds
	case planet == "Jupiter" :
		age = s / jupiteryearsseconds
	case planet == "Saturn" :
		age = s / saturnyearseconds
	case planet == "Uranus" :
		age = s / uranusyearseconds
	case planet == "Neptune" :
		age = s / neptuneyearseconds
		}
	return age
}
