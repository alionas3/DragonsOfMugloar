package gamepkg

import (
	"encoding/json"
	"testing"
)

/*
  This function checks if dragon skills are assigned correctly
*/
func TestCreateDragon(t *testing.T) {
	var (
		game       = Game{}
		gotDragon  = Dragon{}
		wantDragon = Dragon{}
		weather    = Weather{}
	)
	game.GameId = 7569516
	game.Knight.Name = "Sir. Eugene Murphy of Newfoundland and Labrador"
	game.Knight.Attack = 8
	game.Knight.Armor = 3
	game.Knight.Agility = 2
	game.Knight.Endurance = 7
	weather.Code = "NMR"
	weather.Message = "Another day of everyday normal regular weather, business as usual, unless it’s going to be like the time of the Great Paprika Mayonnaise Incident of 2014, that was some pretty nasty stuff."
	dragonJson := CreateDragon(game, weather)
	json.Unmarshal([]byte(dragonJson), &gotDragon)

	wantDragon.Strengths.ScaleThickness = 10
	wantDragon.Strengths.ClawSharpness = 2
	wantDragon.Strengths.WingStrength = 2
	wantDragon.Strengths.FireBreath = 6

	if gotDragon.Strengths.ScaleThickness != wantDragon.Strengths.ScaleThickness ||
		gotDragon.Strengths.ClawSharpness != wantDragon.Strengths.ClawSharpness ||
		gotDragon.Strengths.WingStrength != wantDragon.Strengths.WingStrength ||
		gotDragon.Strengths.FireBreath != wantDragon.Strengths.FireBreath {
		t.Errorf("Dragon skill balancing was incorrect, got: %d, want: %d", gotDragon, wantDragon)
	}
	dragonString := string(dragonJson)
	status, _ := ResolveBattle(dragonString, game.GameId)
	if status != "Victory" {
		t.Errorf("Dragon lost the battle, got: %s, want: %s", status, "Victory")
	}
}
