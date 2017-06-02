package gamepkg

import (
	"os"
	"strconv"
	"io/ioutil"
	"fmt"
)

func LogStatistics(btlWon int, btlLost int,totalBtls int){

	file, error := os.Create("BattleOfMugloarResult.txt")
	checkErr(error)
	defer file.Close()
	result := "Battles won: "+ strconv.Itoa(btlWon)+
		"\r\nBattles lost:"+strconv.Itoa(btlLost)+
		"\r\nWon percentage: "+strconv.Itoa(100-(btlLost * 100)/(totalBtls))+
		"%"
	file.WriteString(result)
	stream,error := ioutil.ReadFile("BattleOfMugloarResult.txt")
	checkErr(error)
	res := string(stream)
	fmt.Println(res)
}
