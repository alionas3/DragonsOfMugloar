package gamepkg

import (
	"encoding/json"
	"fmt"
)

type Dragon struct {
	Strengths struct {
		ScaleThickness int `json:"scaleThickness"`
		ClawSharpness  int `json:"clawSharpness"`
		WingStrength   int `json:"wingStrength"`
		FireBreath     int `json:"fireBreath"`
	} `json:"dragon"`
}

/*
  This function creates dragon json for resolveapi.go
  %param game - GAME, created game from gameapi.go
  %param weather - WEATHER, weather obj from weatherapi.go
  #returns: string - parsed DRAGON json
*/
func CreateDragon(game Game, weather Weather) (dragonString string) {
	dragon := Dragon{}
	var SendDragon string = "Y"
	switch weather.Code {
	//Heavy rain, knights comes with umbrella boats
	case "HVA":
		dragon.Strengths.ScaleThickness = 5
		dragon.Strengths.ClawSharpness = 10
		dragon.Strengths.WingStrength = 5
		dragon.Strengths.FireBreath = 0
	//Long dry, need to balance dragon skills
	case "T E":
		dragon.Strengths.ScaleThickness = 5
		dragon.Strengths.ClawSharpness = 5
		dragon.Strengths.WingStrength = 5
		dragon.Strengths.FireBreath = 5
	//Storm, do not send dragon to battle
	case "SRO":
		SendDragon = "N"
	//Normal weather, normal battles
	case "NMR":
		dragon = getNormalDragon(game)
	//Fog, knights, dragon locating skills reduce
	case "FUNDEFINEDG":
		dragon = getNormalDragon(game)
	}
	dragonJson, error := json.Marshal(dragon)
	dragonString = string(dragonJson)
	checkErr(error)
	//do not send dragon to battle
	if SendDragon == "N" {
		dragonString = ""
	}
	return
}

/*
  This function assigns dragon skill for normal weather
  %param game - GAME, accepts game from GAMEAPI
  #returns: dragon - DRAGON with assigned skills
*/
func getNormalDragon(game Game) (dragon Dragon) {
	dragon = Dragon{}
	var (
		maxKnightKey  string = ""
		minKnightKey  string = ""
		minKnightStat int
		cnt           int = 0
	)
	knightMap := map[string]int{
		"Attack":    game.Knight.Attack,
		"Armor":     game.Knight.Armor,
		"Agility":   game.Knight.Agility,
		"Endurance": game.Knight.Endurance,
	}
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
	switch maxKnightKey {
	case "Armor":
		dragon.Strengths.ClawSharpness = maxKnightStat
	case "Attack":
		dragon.Strengths.ScaleThickness = maxKnightStat
	case "Agility":
		dragon.Strengths.WingStrength = maxKnightStat
	case "Endurance":
		dragon.Strengths.FireBreath = maxKnightStat
	}
	//assigns lowest Knight points to dragon
	for key, value := range knightMap {
		if value != 0 {
			switch key {
			case "Armor":
				dragon.Strengths.ClawSharpness = value - 1
			case "Attack":
				dragon.Strengths.ScaleThickness = value - 1
			case "Agility":
				dragon.Strengths.WingStrength = value - 1
			case "Endurance":
				dragon.Strengths.FireBreath = value - 1
			}
		}
	}
	//assigns left balance to dragon
	switch minKnightKey {
	case "Attack":
		dragon.Strengths.ScaleThickness = 20 - (dragon.Strengths.FireBreath +
			dragon.Strengths.WingStrength +
			dragon.Strengths.ClawSharpness)
	case "Armor":
		dragon.Strengths.ClawSharpness = 20 - (dragon.Strengths.FireBreath +
			dragon.Strengths.WingStrength +
			dragon.Strengths.ScaleThickness)
	case "Agility":
		dragon.Strengths.WingStrength = 20 - (dragon.Strengths.FireBreath +
			dragon.Strengths.ClawSharpness +
			dragon.Strengths.ScaleThickness)
	case "Endurance":
		dragon.Strengths.FireBreath = 20 - (dragon.Strengths.WingStrength +
			dragon.Strengths.ClawSharpness +
			dragon.Strengths.ScaleThickness)

	}

	defer checkDragonPoints(dragon.Strengths.ScaleThickness,
		dragon.Strengths.FireBreath,
		dragon.Strengths.WingStrength,
		dragon.Strengths.ClawSharpness)
	return
}

/*
  This function checks if dragon points are not more than 20
*/
func checkDragonPoints(args ...int) {
	total := 0
	for _, v := range args {
		total += v
	}
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()
	if total > 20 {
		panic("Dragon has more than 20 points!!!")
	}
}
