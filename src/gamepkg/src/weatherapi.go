package gamepkg

import (
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)
type Weather struct{
	Code    string `xml:"code"`
	Message string `xml:"message"`
}

func GetWeather(gameId int)(Weather) {
	weather := Weather{}
	var url string = "http://www.dragonsofmugloar.com/weather/api/report/" + strconv.Itoa(gameId)
	response, error := http.Get(url)
	checkErr(error)
	defer response.Body.Close()
	xmlData, error := ioutil.ReadAll(response.Body)
	xml.Unmarshal([]byte(xmlData),&weather)
	checkErr(error)
	return weather
}
