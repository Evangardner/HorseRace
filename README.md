# HorseRace
  A card game for people who don't have cards at the moment. Written in golang.
## Introduction
  A game where there are four horses on a track, each one corresponding to a unique suit of card. Cards are pulled randomly from a digital 
deck until a horse crosses the finish line! 

## Installation
  ### Prerequisites:
  **golang**: available for download at their [website](https://learn.go.dev/)
  ### MacOS and Linux Installation:
   1. Download Github repository as a .zip file.
   2. Extract to desired location.
   3. Run the following command:
      ```go build /PATHTO.GOFILE```
   4. A UNIX executable will be created and will be ready to run.
   ### Windows Installation:
   1. Download Github repository as a .zip file.
   2. Extract to desired location.
   3. A Microsoft executable is already included in the folder, simply run it however you see fit. 
## Usage
To begin the game, run the executable. A prompt will appear asking you to place your bets. Do so, then press Enter to begin the race. Continue pressing Enter until the race is completed. 

Instruct players to place bets on each horse in the form of seconds spent drinking a 
beverage of their choosing. When a player's horse wins, they may distribute double the number they bet to the rest of the players.

For example, if a player were to bet 3 seconds and their horse won, they would have 6 seconds to distribute to other players. They
could make 3 players drink for 2 seconds, or 2 players drink for 3, or 1 player for 6. A player can give their seconds to themselves
if they wish. 

Players that bet on the winning horse can still receive seconds from the other players who bet on the winning horse. 

Every time all four horses pass a stage of the track, one card is pulled randomly from the deck and the corresponding horse is sent back 1 stage.
## Troubleshooting/Support
Please direct all support inquiries and bug reports to AZ@AZ.net
## FAQs
### Which horse should I place my bet on?
This question is extremely common, and for good reason, as betting on the right horse will win you the game. The best horse to bet on is the one that will win. However, this horse will change between games based on the pseudo random number generator in go. Reverse engineer the randomness algorithm in PRNG to decipher the best horse on which to bet.
### How many people can play the game?
Since the player who wins may distribute seconds to themselves, this game can be played with one player at minimum, and infinitely many players at maximum.
### What is a horse?
A (admittedly imperfect) description of a horse is a mammal with four legs and a mane. Horses are domesticated animals that human beings have bred for the purpose of manual labor and transportation, as well as a variety of other uses. Horses are sometimes bred for the purpose of racing, and are driven around a track by a rider called a "jockey". Prior knowledge of the anatomy, appearance, or behavior of a horse are not required for use of this program. The same can be said for jockeys, which have been omitted from the program entirely. 
## How To Contribute?
Contributions are always welcome to Horse Race. Before committing to the project, please ask permission by emailing our support email (AZ@AZ.net)
## License
[MIT License](https://github.com/Evangardner/HorseRace/blob/master/LICENSE)
