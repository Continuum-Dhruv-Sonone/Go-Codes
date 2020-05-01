package kingdom

var kingdomDetails = make(map[string]string)

// GetEmblem ...
func GetEmblem(name string) string {
	return kingdomDetails[name]

}

// Setup intialise kingdoms
func Setup() {

	kingdoms := []struct {
		name   string
		emblem string
	}{
		{"LAND", "Panda"},
		{"WATER", "Octopus"},
		{"ICE", "Mammoth"},
		{"AIR", "Owl"},
		{"FIRE", "Dragon"},
	}

	for _, kingdom := range kingdoms {
		kingdomDetails[kingdom.name] = kingdom.emblem
	}
}
