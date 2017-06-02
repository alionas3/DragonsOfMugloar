package gamepkg

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Game struct{
	GameId int `json:"gameId"`
	Knight struct {
		       Name   string `json:"name"`
		       Attack    int `json:"attack"`
		       Armor     int `json:"armor"`
		       Agility   int `json:"agility"`
		       Endurance int `json:"endurance"`
	       }`json:"knight"`
}

func GetGame()(Game) {
	game := Game{}
	response, error := http.Get("http://www.dragonsofmugloar.com/api/game")
	checkErr(error)
	defer response.Body.Close()
	jsonData, error := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(jsonData),&game)
	checkErr(error)
	return game
}