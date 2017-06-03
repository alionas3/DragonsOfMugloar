package gamepkg

import (
	"fmt"
	"gamepkg/src"
	"github.com/cheggaaa/pb"
	"bufio"
	"os"
)
/*
  This function start the game
  %params - accepts no parameters
  #returns - does not return anything
*/
func StartBattle() {
	var(
	    totalBtls int = 0
	    btlWon    int = 0
	    btlLost   int = 0
	    dragon    string = ""
	    cnt       int = 0
	)
	fmt.Print("How many battles would you like to play?:")
	_,error := fmt.Scanf("%d", &totalBtls)
	if error != nil {
		fmt.Println("Please enter valid battle number.")
	}
	bar := pb.StartNew(totalBtls)
	bar.ShowBar      = true
	bar.ShowTimeLeft = true
	bar.ShowPercent  = true
	bar.ShowSpeed    = true
	bar.ShowCounters = true
	bar.SetWidth(50)
	bar.Start()
	for cnt < totalBtls {
		//gets the game
		game := gamepkg.GetGame()
		//gets battle weather
		weather := gamepkg.GetWeather(game.GameId)
		//creates dragon
		dragon = gamepkg.CreateDragon(game, weather)
		//resolves battle
		status, _ := gamepkg.ResolveBattle(dragon, game.GameId)
		if status == "Defeat"{
			btlLost++
		}else{
			btlWon++
		}
		cnt++
		bar.Increment()
	}
	bar.FinishPrint("All Battles Finished!")
	gamepkg.LogStatistics(btlWon,btlLost,totalBtls)
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Press ENTER to exit...")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}
}
