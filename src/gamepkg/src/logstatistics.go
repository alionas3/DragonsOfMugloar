package gamepkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

/*
  This function calculates, writes statistics to file and prints the result
  %param btlWon    - int, won battles
  %param btlLost   - int, lost battles
  %param totalBtls - int, total battles
  #returns: does not return anything
*/
func LogStatistics(btlWon int, btlLost int, totalBtls int, elapsed time.Duration) {
	file, error := os.Create(resultFile)
	checkErr(error)
	defer file.Close()
	result := "Battles won: " + strconv.Itoa(btlWon) +
		"\r\nBattles lost:" + strconv.Itoa(btlLost) +
		"\r\nWon percentage: " + strconv.Itoa(100-(btlLost*100)/(totalBtls)) +
		"%" +
		"\r\nElapsed time: " + elapsed.String()
	file.WriteString(result)
	stream, error := ioutil.ReadFile(resultFile)
	checkErr(error)
	res := string(stream)
	fmt.Println(res)
}
