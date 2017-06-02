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

func ResolveBattle(dragon string, gameId int)(string, string){
	result := Result{}
	var url string = "http://www.dragonsofmugloar.com/api/game/"+strconv.Itoa(gameId)+"/solution"
	var jsonDragon = []byte(dragon)
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