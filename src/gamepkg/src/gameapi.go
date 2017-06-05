package gamepkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Game struct {
	GameId int `json:"gameId"`
	Knight struct {
		Name      string `json:"name"`
		Attack    int    `json:"attack"`
		Armor     int    `json:"armor"`
		Agility   int    `json:"agility"`
		Endurance int    `json:"endurance"`
	} `json:"knight"`
}

/*
  This function calls the game api, parses response to json and returns GAME object
  %param: accepts no parameters
  #returns: GAME object
*/
func GetGame() (game Game) {
	game = Game{}
	response, error := http.Get(gameUrl)
	checkErr(error)
	defer response.Body.Close()
	jsonData, error := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(jsonData), &game)
	checkErr(error)
	return
}
