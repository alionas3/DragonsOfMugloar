package gamepkg
/*
  This function handles the error, stops the game if error occurs
  %param err - ERROR, error from other functions
  #returns: does not return anything
*/
func checkErr(err error){
	if err != nil{
		panic(err)
	}
}