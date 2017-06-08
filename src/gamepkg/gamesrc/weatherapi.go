package gamepkg

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Weather struct {
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

func GetWeather(gameId int) (weather Weather) {
	var (
		url = weatherUrl + strconv.Itoa(gameId)
	)
	response, error := http.Get(url)
	checkErr(error)
	defer response.Body.Close()
	xmlData, error := ioutil.ReadAll(response.Body)
	xml.Unmarshal([]byte(xmlData), &weather)
	checkErr(error)
	return
}
