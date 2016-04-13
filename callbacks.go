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
    //fmt.Println(log_type)
//    if log_type == dota.DOTA_COMBATLOG_TYPES(4){
//        count++
//        fmt.Println(count)
//    }
    return nil
}

func ChatEvent(message *dota.CDOTAUserMsg_ChatEvent) error{
    if fmt.Sprintf("%v", message.GetType()) == "CHAT_MESSAGE_HERO_KILL" {
        count++
        fmt.Println(count)
    }
    //fmt.Println(message.GetType())
    return nil
}

func UnitEvent(message *dota.CDOTAUserMsg_UnitEvent) error{
    //fmt.Println(message["msg_type"])
    return nil
}