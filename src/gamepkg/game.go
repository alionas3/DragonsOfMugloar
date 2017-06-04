package gamepkg

import (
	"fmt"
	"gamepkg/src"
	"github.com/cheggaaa/pb"
	"bufio"
	"os"
	"sync"
	"time"
)
/*
  This function start the game
  %params - accepts no parameters
  #returns - does not return anything
*/
func StartBattle() {
	var(
	    totalBtls int = 0
	    btlLost   int = 0
	    dragon    string = ""
	    cnt       int = 0
	    btlWon    int = 0
	    wg        sync.WaitGroup
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
	start := time.Now()
	bar.Start()
	for i := 0; i < totalBtls; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//gets the game
			game := gamepkg.GetGame()
			//gets battle weather
			weather := gamepkg.GetWeather(game.GameId)
			fmt.Println(weather)
			//creates dragon
			dragon = gamepkg.CreateDragon(game, weather)
			//resolves battle
			status, _ := gamepkg.ResolveBattle(dragon, game.GameId)
			if status == "Defeat" {
				btlLost++
			}else{
				btlWon++
			}
			cnt++
			bar.Increment()
		}()
	}
	wg.Wait()//waits for all go routines to finish
	bar.FinishPrint("All Battles Finished!")
	elapsed := time.Since(start)//gets the elapsed time of go routines
	gamepkg.LogStatistics(btlWon,btlLost,totalBtls,elapsed)
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Press ENTER to exit...")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}
}