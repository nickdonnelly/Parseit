package main

import(
    "sort"
)

func GetPrintableStringFromVariableName(varName string) string{
    switch varName{
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


func GetAlphabetizedKeyListFromMap(myMap map[string]int) []string{
    var keys []string
    for k := range myMap{
        keys = append(keys, k)
    }
    sort.Strings(keys)
    return keys
}