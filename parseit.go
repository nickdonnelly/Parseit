package main

import(
    //"log"
    "fmt"
    "flag"
    "os"
    //"github.com/dotabuff/manta"
    //"github.com/dotabuff/manta/dota"
)


func main(){
    
    var modePtr = flag.String("mode", "text", "The mode of output for this match.")
    var matchPtr = flag.String("match", "", "The match ID you would like to parse")
    //var offlinePtr = flag.Bool("offline", false, "Set this to true if you want the match to be downloaded directly instead of found in the current directory.")
    var heroPtr = flag.String("hero", "", "Use this to select the hero you would like to parse data for (only for image mode). Use full names, all lowercase. Separate separate words using a dash character (-).")
    
    flag.Parse()
    
    if *matchPtr == ""{
        fmt.Println("No match ID provided. Try again using the --match=match_id_here flag.")
        os.Exit(2) // code 2 is match-not-found
    }
    
    switch *modePtr{
    case "text":
        textParse()
    case "image":
        if *heroPtr == ""{
            fmt.Println("You selected image mode but didn't provide a hero. Please include the --hero=hero_name flag and try again.")
        }else{
            imageParse()
        }
    default:
        fmt.Println("Mode was set incorrectly. Possible options: \ntext \nimage")
        os.Exit(2)
    }
    
}

func textParse(){
    
}

func imageParse(){
    
}