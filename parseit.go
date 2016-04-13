package main

import(
    //"log"
    "fmt"
    "flag"
    "os"
    "github.com/dotabuff/manta"
)    

var modePtr = flag.String("mode", "text", "The mode of output for this match.")
var matchPtr = flag.String("match", "", "The match file you would like to parse")
//var offlinePtr = flag.Bool("offline", false, "Set this to true if you want the match to be downloaded directly instead of found in the current directory.")
var heroPtr = flag.String("hero", "", "Use this to select the hero you would like to parse data for (only for image mode). Use full names, all lowercase. Separate separate words using a dash character (-).")


var heroes [10]Hero

type Hero struct{
    name string
    kills int
    deaths int
    assists int
    gold int
    networth int
}


func main(){

    
    flag.Parse()
    
    if *matchPtr == ""{
        fmt.Println("No match file provided. Try again using the --match=match_file_here flag.")
        os.Exit(2) // code 2 is match-not-found
    }
    
    switch *modePtr{
    case "text":
        textParse()
    case "image":
        if *heroPtr == ""{
            fmt.Println("You selected image mode but didn't provide a hero. Please include the --hero=hero_name flag and try again.")
            os.Exit(2)
        }else{
            imageParse()
        }
    default:
        fmt.Println("Mode was set incorrectly. Possible options: \ntext \nimage")
        os.Exit(2)
    }
    
}

func textParse(){
    parser, err := manta.NewParserFromFile(*matchPtr)
    if err != nil{
        fmt.Printf("Unable to create match parser: %s", err)
        os.Exit(5)
    }
    parser.Callbacks.OnCUserMessageSayText2(AllchatMessage)
    parser.Callbacks.OnCDOTAClientMsg_MatchMetadata(MatchMetadata)
    parser.Callbacks.OnCMsgDOTACombatLogEntry(CombatLogMessage)
    parser.Callbacks.OnCDOTAUserMsg_ChatEvent(ChatEvent)
    parser.Callbacks.OnCDOTAUserMsg_UnitEvent(UnitEvent)
    
    parser.Start()
    fmt.Println("Parsing complete. Printing report.")
    printTextReport()
}

func imageParse(){
    
}



func printTextReport(){
    
}