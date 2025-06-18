package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avengers: Endgame"
	case "tt1825683":
		return "Black Panther"
	case "tt2250912":
		return "Good Old Days"

	}
	return "not found."
}
