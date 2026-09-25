package main

import "math/rand/v2"

type Name struct{ First, Last string }
type namePool struct{ first, prefix, suffix []string }

var ancestries = []string{"Human", "Orc", "Night elf", "Dwarf", "Tauren", "Blood elf", "Undead"}
var pools = map[string]namePool{
	"Human":     {[]string{"Alden", "Mira", "Rowan", "Elara", "Cedric", "Tessa", "Brenna", "Galen"}, []string{"Ash", "Bright", "West", "Raven", "Fair", "Storm"}, []string{"ford", "wood", "vale", "brook", "field", "well"}},
	"Orc":       {[]string{"Gorak", "Draka", "Thurg", "Rokka", "Mogra", "Karg", "Urza", "Durok"}, []string{"Iron", "Skull", "Blood", "Stone", "Wolf", "Thunder"}, []string{"fist", "breaker", "fang", "scar", "axe", "howl"}},
	"Night elf": {[]string{"Ilyra", "Thalen", "Seranis", "Velora", "Aelrin", "Naelis", "Luneth", "Vaelin"}, []string{"Moon", "Star", "Dusk", "Shadow", "Silver", "Night"}, []string{"whisper", "bloom", "leaf", "song", "weaver", "bough"}},
	"Dwarf":     {[]string{"Borin", "Helga", "Dagna", "Thorin", "Brenna", "Kelda", "Durin", "Orik"}, []string{"Bronze", "Iron", "Deep", "Ember", "Stone", "Ale"}, []string{"beard", "hammer", "anvil", "forge", "delver", "mantle"}},
	"Tauren":    {[]string{"Hamu", "Tala", "Maku", "Nara", "Orun", "Aponi", "Kova", "Tahu"}, []string{"Dawn", "Cloud", "River", "Earth", "Wind", "Sun"}, []string{"hoof", "mane", "walker", "horn", "runner", "seeker"}},
	"Blood elf": {[]string{"Aethas", "Lyriel", "Vaelis", "Selara", "Theron", "Faelyn", "Aeris", "Belara"}, []string{"Sun", "Dawn", "Flame", "Gold", "Bright", "Ember"}, []string{"strider", "song", "weaver", "gaze", "heart", "wing"}},
	"Undead":    {[]string{"Mordren", "Vesra", "Ravik", "Sylith", "Keldran", "Morwen", "Thessik", "Velric"}, []string{"Grave", "Dusk", "Bone", "Ash", "Blight", "Raven"}, []string{"whisper", "veil", "mourn", "shroud", "binder", "shade"}},
}

func generate(ancestry string, count int) ([]Name, bool) {
	pool, ok := pools[ancestry]
	if !ok || count < 1 || count > len(pool.first)*len(pool.prefix)*len(pool.suffix) {
		return nil, false
	}
	names := make([]Name, 0, count)
	seen := make(map[Name]bool)
	pick := func(values []string) string { return values[rand.IntN(len(values))] }
	for len(names) < count {
		name := Name{pick(pool.first), pick(pool.prefix) + pick(pool.suffix)}
		if !seen[name] {
			names = append(names, name)
			seen[name] = true
		}
	}
	return names, true
}
