# HorseRace
  a card game for people who don't have cards at the moment. written in golang.
## Description
  A game where there are four horses on a track, each one corresponding to a suit of card. Cards are pulled randomly from a digital 
deck until the horse crosses the finish line! 

## Installation
  ### Prerequisites:
  **golang**: available for download at their [website](https://learn.go.dev/)
  ### MacOS and Linux Installation:
   1. Download Github repository as a .zip file
   2. Extract to desired location
   3. run the following command:
      ```go build /PATHTO.GOFILE```
   4. A unix executable will be created and will be ready to run.
   ### Windows Installation:
   1. Download Github repository as a .zip file
   2. Extract to desired location
   3. A Microsoft executable is already included in the folder, simply run it however you see fit. 
## Usage
To begin the game, run the executable. A prompt will appear asking you to place your bets. Do so, then press Enter to begin the race. Continue pressing enter until the race is completed. 

Players place bets on each horse, usually in the form of seconds spent drinking a 
beverage of their choosing. If a player's horse wins, they may distribute double the number they bet to the rest of the players.

For example, if a player were to bet 3 seconds and their horse won, they would have 6 seconds to distribute to other players. They
could make 3 players drink for 2 seconds, or 2 players drink for 3, or 1 player for 6. A player can give their seconds to themselves
if they wish. 

Players that also bet on the winning horse can still receive seconds from the other players who won. 

Every time all four horses pass a stage of the track, one card is pulled randomly from the deck and the corresponding horse is sent back.
## License
[MIT License](https://github.com/Evangardner/HorseRace/blob/master/LICENSE)
