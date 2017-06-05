## Dragons Of Mugloar
This is a solution for the game Dragons of Mugloar which can be found at www.dragonsofmugloar.com. This game is implemented with GO as a terminal based program.

<h3><b>What is the game about?</b></h3>

This game is about medieval times when the dragons where still alive and knights so brave, that they risk their lives to prove their courage in the field of battle. Everyday knights take a brave journey to conquer the kingdom and take the indescribably gorgeous princesses for any cost.

<h3><b>What is the purpose of this game?</b></h3>

You play as a Head of Dragon Resource Management and as you may know, knights always trying to attack your kingdom and take the princess, so your main goal is to keep your kingdom and dragon(especially dragons) always alive by selectively assigning 20 points to 4 skills for your dragon which are:
```javascript
Scale thickness
Claw sharpness
Wing strength
Fire breath
```
<h3><b>Game Marks:</b></h3>

Before entering the battle you should know that the weather in our kingdom can be pretty hectic, so keep in mind couple of things:
<br />**Normal weather** - normal fights
<br />**Storm** - everyone dies
<br />**Heavy rain with floods** - knights come by boats with umbrella, so the fire is useless in the rain and additional claw sharpening is needed to destroy the boats with umbrellas
<br />**The long dry** - Only the ZENNEST dragons who have achieved great balance of their inner-self through meditation can win battles during the drought!
<br />**Fog** - All dragons come with excellent knight-locating skills, no puny fog can change that. The knights dragon-locating skills, however are limited to understanding whether or not they have already been eaten.

<h3><b>What should I do to start the game?</b></h3>

Firstly you should install GO into your machine, click this [link](https://golang.org/doc/install) to do that.After instalation is completed, open your cmd or terminal and execute the following commands:

```javascript
go get github.com/alionas3/DragonsOfMugloar
go build DragonsOfMugloar
go run src\dragonsofmugloar\main.go
```

<h3><b>How should i play the game?</b></h3>

After launching the game, it will ask you a question, how much battles do you wish to fight?. At his point you should enter a whole number, as you can see I've enter 200 battles, after that hit ENTER.

![Alt text](http://i.imgur.com/RnLV7L8.png)

After all battles are finnished, you will receive an output:

![Alt text](http://i.imgur.com/xDTV1sc.png)

*progress bar - shows how much battles left<br>
*Battles won - shows how much battles did you managed to win<br>
*Battles lost - shows how much battles did you lost<br>
*Won percentage - shows how much battles did you won in percents<br>
*Elapsed time - shows how much time did it take to fight all the battles (at this point it took 1,7 seconds)

This output is saved in text file *BattleOfMugloarResult.txt* which could be found in your game directory.

<h3><b>Errors</b></h3>

If you will enter something else, but not a number(when the program is launched for the first time) you will receive an error and program will exit.

![Alt text](http://i.imgur.com/bAdzd2c.png)







