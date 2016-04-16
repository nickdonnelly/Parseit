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

func MatchMetadata(message *dota.CDOTAClientMsg_MatchMetadata) error {
	//fmt.Println("Got match metadata.");
	return nil
}

var count = 0

func CombatLogMessage(message *dota.CMsgDOTACombatLogEntry) error {
  if message.GetIsAttackerHero() && message.GetIsTargetHero() && uint32(message.GetType()) == 4{
    count++
    if count == 7{
      fmt.Println("Ability lvl", message.GetAbilityLevel())
      fmt.Println("Assist 0", message.GetAssistPlayer0())
      fmt.Println("Assist n", message.GetAssistPlayer1())
      fmt.Println("Assist n", message.GetAssistPlayer2())
      fmt.Println("Assist n", message.GetAssistPlayer3())
      fmt.Println("Assist players", message.GetAssistPlayers())
      fmt.Println("Hero lvl", message.GetAttackerHeroLevel())
      fmt.Println("Attacker name", message.GetAttackerName())
      fmt.Println("Attacker TEam", message.GetAttackerTeam())
      fmt.Println("Building type", message.GetBuildingType())
      fmt.Println("dmg category", message.GetDamageCategory())
      fmt.Println("event location", message.GetEventLocation())
      fmt.Println("damg type", message.GetDamageType())
      fmt.Println("gpm", message.GetGpm())
      fmt.Println("xpm", message.GetXpm())
      fmt.Println("xp reason", message.GetXpReason())
      fmt.Println("lasthits", message.GetLastHits())
      fmt.Println("timestamp", message.GetTimestamp())
      fmt.Println("timestamp raw", message.GetTimestampRaw())
      fmt.Println("target team", message.GetTargetTeam())
      fmt.Println("target name", message.GetTargetName())
      fmt.Println("target source name", message.GetTargetIsSelf())
    }
  }
  // if message.GetIsTargetHero() && uint32(message.GetType()) == 4{
  //   if GetHeroStringById(message.GetAttackerName()) == "Lion"{
  //     count++
  //     fmt.Println(count, " - ", message.GetInflictorName())
  //   }
  //   fmt.Println("Killer: ", GetHeroStringById(message.GetAttackerName()))
  // }
	// if message.GetIsTargetHero() && message.GetType() == 4 {
    // fmt.Println(GetHeroStringById(message.GetTargetSourceName()), message.GetTargetHeroLevel())
		// fmt.Println("A levelh ", message.GetAttackerHeroLevel(), GetHeroStringById(message.GetAttackerName()), "killed a level", message.GetTargetHeroLevel(), message.GetTargetName())
	// }
	return nil
}

func ChatEvent(message *dota.CDOTAUserMsg_ChatEvent) error {
	mType := fmt.Sprintf("%v", message.GetType())
	switch mType {
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

func UnitEvent(message *dota.CDOTAUserMsg_UnitEvent) error {
	//fmt.Println(message["msg_type"])
	return nil
}

func CombatLogEvent(message *manta.GameEvent) error {
	fmt.Println("Combat log event grabbed!")
	return nil
}

func InitChatResultsMap() {
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
