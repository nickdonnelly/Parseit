package main

import (
    "fmt"
    "github.com/dotabuff/manta"
    "github.com/dotabuff/manta/dota"
)


var ChatResult = make(map[string]int)

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
    if message.GetIsTargetHero() && message.GetType() == 4 {
      fmt.Println("Hero died!")
    }
    return nil
}

func ChatEvent(message *dota.CDOTAUserMsg_ChatEvent) error{
    mType := fmt.Sprintf("%v", message.GetType())
    switch mType{
        case "CHAT_MESSAGE_HERO_KILL":
            fallthrough
        case "CHAT_MESSAGE_HERO_DENY":
            ChatResult["killcount"]++
        case "CHAT_MESSAGE_TOWER_KILL":
            fallthrough
        case "CHAT_MESSAGE_TOWER_DENY":
            ChatResult["towerkills"]++
        case "CHAT_MESSAGE_RUNE_PICKUP":
            ChatResult["runestaken"]++
        case "CHAT_MESSAGE_BUYBACK":
            ChatResult["buybacks"]++
        case "CHAT_MESSAGE_DISCONNECT":
            ChatResult["disconnects"]++
        case "CHAT_MESSAGE_STREAK_KILL":
            ChatResult["streakkills"]++
        case "CHAT_MESSAGE_GLYPH_USED":
            ChatResult["glyphsused"]++
        case "CHAT_MESSAGE_RANDOM":
            ChatResult["heroesrandomed"]++
        case "CHAT_MESSAGE_PAUSED":
            ChatResult["pauses"]++
        case "CHAT_MESSAGE_EFFIGY_KILL":
            ChatResult["effigykills"]++
        case "CHAT_MESSAGE_DENIED_AEGIS":
            ChatResult["aegiesdenies"]++
        case "CHAT_MESSAGE_AEGIS":
            ChatResult["aegispickups"]++
        case "CHAT_MESSAGE_AEGIS_STOLEN":
            ChatResult["aegissteals"]++
        case "CHAT_MESSAGE_COURIER_LOST":
            ChatResult["courierdeaths"]++
        case "CHAT_MESSAGE_ITEM_PURCHASE":
            ChatResult["itemspurchased"]++
        case "CHAT_MESSAGE_VICTORY_PREDICTION_SINGLE_USER_CONFIRM":
            ChatResult["victorypredictions"]++

    }
    return nil
}

func UnitEvent(message *dota.CDOTAUserMsg_UnitEvent) error{
    //fmt.Println(message["msg_type"])
    return nil
}

func CombatLogEvent(message *manta.GameEvent) error{
  fmt.Println("Combat log event grabbed!")
  return nil
}

func InitChatResultsMap(){
    ChatResult["killcount"] = 0
    ChatResult["towerkills"] = 0
    ChatResult["runestaken"] = 0
    ChatResult["buybacks"] = 0
    ChatResult["disconnects"] = 0
    ChatResult["streakkills"] = 0
    ChatResult["glyphsused"] = 0
    ChatResult["heroesrandomed"] = 0
    ChatResult["pauses"] = 0
    ChatResult["effigykills"] = 0
    ChatResult["aegisdenies"] = 0
    ChatResult["aegispickups"] = 0
    ChatResult["aegissteals"] = 0
    ChatResult["courierdeaths"] = 0
    ChatResult["itemspurchased"] = 0
    ChatResult["victorypredictions"] = 0
}
