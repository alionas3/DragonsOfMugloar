package gamepkg

import (
	"strconv"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)
type Result struct{
	Status  string `json:"status"`
	Message string `json:"message"`
}
/*
  This function calls solution api and solves the battle
  %param dragon - string, parsed dragon json
  %param gameId - int, game code from gameapi
  #returns: Status - string, status of the resolved bttle
  	    Message - string, message of the resolved battle
*/
func ResolveBattle(dragon string, gameId int)(string, string){
	var(
	    result     = Result{}
	    url        = "http://www.dragonsofmugloar.com/api/game/"+strconv.Itoa(gameId)+"/solution"
	    jsonDragon = []byte(dragon)
	)
	req, error := http.NewRequest("PUT", url,bytes.NewBuffer(jsonDragon))
	checkErr(error)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, error := client.Do(req)
	checkErr(error)
	defer resp.Body.Close()
	jsonData, error := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(jsonData),&result)
	checkErr(error)
	return result.Status, result.Message
}