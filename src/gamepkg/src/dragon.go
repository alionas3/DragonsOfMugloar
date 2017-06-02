package gamepkg

import "encoding/json"

type Dragon struct{
	Strengths struct {
			  ScaleThickness int `json:"scaleThickness"`
			  ClawSharpness  int `json:"clawSharpness"`
			  WingStrength   int `json:"wingStrength"`
			  FireBreath     int `json:"fireBreath"`
		  }`json:"dragon"`
}
/*
  This function creates dragon json for resolveapi.go
  %param game - GAME, created game from gameapi.go
  %param weather - WEATHER, weather obj from weatherapi.go
  #returns: string - parsed DRAGON json
*/
func CreateDragon(game Game, weather Weather)(string) {
	dragon := Dragon{}
	var SendDragon   string = "Y"
	//Heavy rain, knights comes with umbrella boats
	if weather.Code == "HVA"{
		dragon.Strengths.ScaleThickness = 5
		dragon.Strengths.ClawSharpness  = 10
		dragon.Strengths.WingStrength   = 5
		dragon.Strengths.FireBreath     = 0
		//Long dry, need to balance dragon skills
	}else if weather.Code == "T E"{
		dragon.Strengths.ScaleThickness = 5
		dragon.Strengths.ClawSharpness  = 5
		dragon.Strengths.WingStrength   = 5
		dragon.Strengths.FireBreath     = 5
		//Storm, do not send dragon to battle
	}else if weather.Code == "SRO"{
		SendDragon = "N"
		//Normal weather, normal battles
	}else if weather.Code == "NMR" || weather.Code == "FUNDEFINEDG"{
		dragon = getNormalDragon(game)
	}
	dragonJson, error := json.Marshal(dragon)
	dragonString := string(dragonJson)
	checkErr(error)
	//do not send dragon to battle
	if SendDragon == "N"{
		dragonString = ""
	}
	return dragonString
}
/*
  This function assigns dragon skill for normal weather
  %param game - GAME, accepts game from GAMEAPI
  #returns: dragon - DRAGON with assigned skills
*/
func getNormalDragon(game Game)(Dragon) {
	dragon := Dragon{}
	var maxKnightKey string = ""
	var minKnightKey string = ""
	var minKnightStat int
	var cnt 	  int = 0
	knightMap := make(map[string]int)
	knightMap["Attack"]    = game.Knight.Attack
	knightMap["Armor"]     = game.Knight.Armor
	knightMap["Agility"]   = game.Knight.Agility
	knightMap["Endurance"] = game.Knight.Endurance
	//assume that the first value is the smallest
	maxKnightStat := knightMap["Armor"]
	//gets maximum Knight skill and minimum Knight skill
	for key, value := range knightMap {
		if value >= maxKnightStat {
			maxKnightStat = value
			maxKnightKey = key
		}
	}
	//assigns +2 to max knight stat
	maxKnightStat = maxKnightStat + 2
	//deletes max Knight skill from map
	delete(knightMap, maxKnightKey)
	for key, value := range knightMap {

		if cnt == 0 {
			minKnightStat = value
		}
		if value <= minKnightStat {
			minKnightKey = key
		}
		cnt = cnt + 1
	}
	//assigns max Knight stat to dragon
	if maxKnightKey == "Armor" {
		dragon.Strengths.ClawSharpness 	= maxKnightStat
	} else if maxKnightKey == "Attack" {
		dragon.Strengths.ScaleThickness = maxKnightStat
	} else if maxKnightKey == "Agility" {
		dragon.Strengths.WingStrength 	= maxKnightStat
	} else if maxKnightKey == "Endurance" {
		dragon.Strengths.FireBreath 	= maxKnightStat
	}
	//assigns lowest Knight points to dragon
	for key, value := range knightMap {
		if value != 0 {
			if key == "Armor" {
				dragon.Strengths.ClawSharpness = value - 1
			} else if key == "Attack" {
				dragon.Strengths.ScaleThickness = value - 1
			} else if key == "Agility" {
				dragon.Strengths.WingStrength = value - 1
			} else if key == "Endurance" {
				dragon.Strengths.FireBreath = value - 1
			}
		}
	}
	if minKnightKey == "Attack" {
		dragon.Strengths.ScaleThickness = 20 - (dragon.Strengths.FireBreath   +
			dragon.Strengths.WingStrength +
			dragon.Strengths.ClawSharpness)
	} else if minKnightKey == "Armor" {
		dragon.Strengths.ClawSharpness  = 20 - (dragon.Strengths.FireBreath    +
			dragon.Strengths.WingStrength  +
			dragon.Strengths.ScaleThickness)

	} else if minKnightKey == "Agility" {
		dragon.Strengths.WingStrength   = 20 - (dragon.Strengths.FireBreath    +
			dragon.Strengths.ClawSharpness +
			dragon.Strengths.ScaleThickness)

	} else if minKnightKey == "Endurance" {
		dragon.Strengths.FireBreath     = 20 - (dragon.Strengths.WingStrength  +
			dragon.Strengths.ClawSharpness +
			dragon.Strengths.ScaleThickness)
	}
	return dragon
}
