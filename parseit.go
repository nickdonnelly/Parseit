package main

import(
    "fmt"
    "flag"
    "os"
)

func main(){
    
    //var modePtr = flag.String("mode", "text", "The mode of output for this match.")
    var matchPtr = flag.String("match", "", "The match ID you would like to parse")
    //var offlinePtr = flag.Bool("offline", false, "Set this to true if you want the match to be downloaded directly instead of found in the current directory.")
    
    flag.Parse()
    
    if *matchPtr == ""{
        fmt.Println("No match ID provided. Try again using the --match=match_id_here flag.")
        os.Exit(2) // code 2 is match-not-found
    }
    
    
}