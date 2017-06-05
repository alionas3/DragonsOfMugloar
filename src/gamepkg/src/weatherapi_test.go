package gamepkg

import "testing"
/*
This function will test weather api
*/


func TestGetWeather(t *testing.T) {
	var (
		wantWeather string = "NMR"
		gotWeather         = Weather{}
		gameId      int    = 3792829
	)
	gotWeather = GetWeather(gameId)

	if gotWeather.Code != wantWeather {
		t.Errorf("Weather is incorrect want: %s, got: %s", wantWeather, gotWeather.Code)
	}
}
