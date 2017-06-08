package gamepkg

import (
	"gamepkg/gamesrc"
	"testing"
)

/*
This function will test weather api
*/

func TestGetWeather(t *testing.T) {
	var (
		wantWeather string = "NMR"
		gotWeather         = gamepkg.Weather{}
		gameId      int    = 3792829
	)
	gotWeather = gamepkg.GetWeather(gameId)

	if gotWeather.Code != wantWeather {
		t.Errorf("Weather is incorrect want: %s, got: %s", wantWeather, gotWeather.Code)
	}
}
