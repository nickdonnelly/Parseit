package main

import (
	//"log"
	"bufio"
	"flag"
	"fmt"
	"github.com/dotabuff/manta"
	"os"
	"strings"
)

var modePtr = flag.String("mode", "text", "The mode of output for this match.")
var matchPtr = flag.String("match", "", "The match file you would like to parse")

//var offlinePtr = flag.Bool("offline", false, "Set this to true if you want the match to be downloaded directly instead of found in the current directory.")
var heroPtr = flag.String("hero", "", "Use this to select the hero you would like to parse data for (only for image mode). Use full names, all lowercase. Separate separate words using a dash character (-).")
var outputFilePtr = flag.Bool("save-external", false, "Use this to save the text report as an external file instead of only printing it to console.")
var outputFilenamePtr = flag.String("output-file", "", "This flag specifies a custom filename for text reports. By default they will simply be named with the match ID. Include a file extension.")

// var heroes [10]Hero

func main() {

	flag.Parse()

	if *matchPtr == "" {
		fmt.Println("No match file provided. Try again using the --match=match_file_here flag.")
		os.Exit(2) // code 2 is match-not-found
	}

	switch *modePtr {
	case "text":
		textParse()
		if *outputFilePtr && *outputFilenamePtr == "" {
			saveTextReport(strings.TrimSuffix(*matchPtr, ".dem") + ".txt")
		} else if *outputFilePtr {
			saveTextReport(*outputFilenamePtr)
		}
	case "image":
		if *heroPtr == "" {
			fmt.Println("You selected image mode but didn't provide a hero. Please include the --hero=hero_name flag and try again.")
			os.Exit(2)
		} else {
			imageParse()
		}
	default:
		fmt.Println("Mode was set incorrectly. Possible options: \ntext \nimage")
		os.Exit(2)
	}

}

func textParse() {
	parser, err := manta.NewParserFromFile(*matchPtr)
	if err != nil {
		fmt.Printf("Unable to create match parser: %s", err)
		os.Exit(5)
	}
	parser.Callbacks.OnCUserMessageSayText2(AllchatMessage)
	parser.Callbacks.OnCDOTAClientMsg_MatchMetadata(MatchMetadata)
	parser.Callbacks.OnCMsgDOTACombatLogEntry(CombatLogMessage)
	parser.Callbacks.OnCDOTAUserMsg_ChatEvent(ChatEvent)
	parser.Callbacks.OnCDOTAUserMsg_UnitEvent(UnitEvent)
	parser.Callbacks.OnCSVCMsg_CreateStringTable(CreateStringTable)
	// parser.Callbacks.OnCSVCMsg_UpdateStringTable(UpdateStringTable)
	parser.OnGameEvent("dota_combatlog", CombatLogEvent)

	fmt.Println("Initializing count structures...")
	InitResultsMaps() // This allows 0s to be reported, i.e. if there are 0 disconnects and you don't have this line, then disconnects will be completely ommitted from the report.
	fmt.Println("Starting parser...")
	parser.Start()
	fmt.Println(parser.LookupStringByIndex("CombatLogNames", 	1))
	fmt.Println(parser.LookupStringByIndex("CombatLogNames", 2))
	fmt.Println(parser.LookupStringByIndex("CombatLogNames", 11))
	fmt.Println(parser.LookupStringByIndex("CombatLogNames", 13))
	fmt.Println("Parsing complete. Printing report.")
	printTextReport(parser)
}

func imageParse() {

}

func printTextReport(parser *manta.Parser) {
	chatResultKeys := GetAlphabetizedKeyListFromMap(ChatResult)
	fmt.Println("+---------------------------------+")
	for k := range chatResultKeys {
		s := "| "
		s += fmt.Sprintf("%24v: ", GetPrintableStringFromVariableName(chatResultKeys[k]))
		s += fmt.Sprintf("%05v |\n", ChatResult[chatResultKeys[k]])
		fmt.Println(s)
	}

	for k, v := range HeroDeaths{
		heroName, g := parser.LookupStringByIndex("CombatLogNames", int32(k))
		fmt.Println(g, heroName, " died ", v, " times.")
	}
	fmt.Println("+---------------------------------+")
}

func saveTextReport(filename string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()
	writer := bufio.NewWriter(f)
	chatResultKeys := GetAlphabetizedKeyListFromMap(ChatResult)
	writer.WriteString("+---------------------------------+\n")
	for k := range chatResultKeys {
		s := "| "
		s += fmt.Sprintf("%24v: ", GetPrintableStringFromVariableName(chatResultKeys[k]))
		s += fmt.Sprintf("%05v |\n", ChatResult[chatResultKeys[k]])
		writer.WriteString(s)
	}
	writer.WriteString("+---------------------------------+")
	f.Sync()
	writer.Flush()
	fmt.Println("Printed report into " + filename + ".")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
