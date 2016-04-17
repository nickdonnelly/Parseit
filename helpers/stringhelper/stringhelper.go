package stringhelper

import (
	"sort"
	"strings"
	"parseit/gamedata/gameconstants"
)

func GetPrintableStringFromVariableName(varName string) string {
	switch varName {
	case "killcount":
		return "Total hero kills"
	case "towerkills":
		return "Total tower kills"
	case "runestaken":
		return "Total runes taken"
	case "buybacks":
		return "Total buybacks"
	case "disconnects":
		return "Total disconnects"
	case "streakkills":
		return "Streak kills"
	case "glyphsused":
		return "Glyphs used"
	case "heroesrandomed":
		return "Heroes randomed"
	case "pauses":
		return "Pauses"
	case "effigykills":
		return "Effigies killed"
	case "aegisdenies":
		return "Aegises denied"
	case "aegispickups":
		return "Aegises picked up"
	case "aegissteals":
		return "Aegises stolen"
	case "courierdeaths":
		return "Courier deaths"
	case "itemspurchased":
		return "Major items purchased"
	case "victorypredictions":
		return "Victory predictions"
	default:
		return varName
	}
}

// Returns the printable hero name for a given hero id. Stored in a literal map, however instead of returning 0 values for unknown keys, it returns the string 'Unknown Hero'
func GetHeroStringById(heroId uint32) string {
	if gameconstants.HeroNames[heroId] == "" {
		return "Unknown Hero"
	}else{
		return gameconstants.HeroNames[heroId]
	}
}

// Returns the printable hero name for a given internal hero name.
func GetHeroStringByInternalName(internalName string) string {
	trimmedString := strings.TrimPrefix(internalName, "npc_dota_hero_")
	if gameconstants.HeroInternalNames[trimmedString] == "" {
		return "Unknown Hero"
	}else{
		return gameconstants.HeroInternalNames[trimmedString]
	}
}

func GetAlphabetizedKeyListFromMap(myMap map[string]string) []string {
	var keys []string
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
