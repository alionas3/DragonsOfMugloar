package gamepkg

import (
	"fmt"
	"gamepkg/src"
	"sync"
	"time"
	"github.com/cheggaaa/pb"
)

/*
  This function start the game
  %params - accepts no parameters
  #returns - does not return anything
*/
func StartBattle() {
	var (
		totalBtls int
		btlLost   int
		btlWon    int
		wg        sync.WaitGroup
		resultChannel chan string = make(chan string)
	)
	fmt.Print("How many battles would you like to play?:")
	_, error := fmt.Scanf("%d", &totalBtls)
	if error != nil {
		fmt.Println("Please enter a valid battle number, program now will exit...")
		return
	}
	bar := pb.StartNew(totalBtls)
	bar.ShowBar = true
	bar.ShowTimeLeft = true
	bar.ShowPercent = true
	bar.ShowSpeed = true
	bar.ShowCounters = true
	bar.SetWidth(45)
	bar.Start()
	start := time.Now()
	for i := 0; i < totalBtls; i++ {
		wg.Add(1)
		go func(resultChannel chan string) {
			//gets the game
			game := gamepkg.GetGame()
			//gets battle weather
			weather := gamepkg.GetWeather(game.GameId)
			//creates dragon
			dragon := gamepkg.CreateDragon(game, weather)
			//resolves battle
			status,_ := gamepkg.ResolveBattle(dragon, game.GameId)
			//assigns the game status to result channel
			resultChannel <- status
		}(resultChannel)

		go func(resultChannel chan string) {
			defer wg.Done()
			status := <-resultChannel
			switch status {
			case "Defeat":
				btlLost++
			case "Victory":
				btlWon++
			}
			bar.Increment()
		}(resultChannel)
		time.Sleep(time.Millisecond * 2) //adding some sleep, because of the server which is failing to handle 1000 requests at the same time
	}
	wg.Wait() //waits for all go routines to finish
	bar.FinishPrint("All Battles Finished!")
	elapsed := time.Since(start) //gets the elapsed time of go routines
	gamepkg.LogStatistics(btlWon, btlLost, totalBtls, elapsed)
}