package main

import (
    "fmt"
    //"github.com/dotabuff/manta"
    "github.com/dotabuff/manta/dota"
)

var count = 0


func AllchatMessage(m *dota.CUserMessageSayText2) error {
    //fmt.Printf("%s said: %s\n", m.GetParam1(), m.GetMessagename())
    return nil
}

func MatchMetadata(message *dota.CDOTAClientMsg_MatchMetadata) error{
    //fmt.Println("Got match metadata.");
    return nil
}

func CombatLogMessage(message *dota.CMsgDOTACombatLogEntry) error{
    log_type := message.GetType()
    fmt.Println(log_type)
    
    return nil
}

func UnitEvent(message *dota.CDOTAUserMsg_UnitEvent) error{
    //fmt.Println(message["msg_type"])
    return nil
}