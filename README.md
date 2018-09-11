Requirements:

 - The computer randomly removes 1 or 2 sticks, but cannot remove more sticks than are left.
 - The human is prompted at each turn for how many sticks he or she wants to remove.
 - Don't accept the user's input if it is illegal. Continue prompting until you get a valid input.
 - Your program must continue play until somebody (computer or human) wins.

On each turn, a player must remove either 1 or 2 sticks from the pile. The goal of the game is to be the player who removes the last stick.
 desired output
Computer starts

Program output:

```
 Round 0, 7 sticks at start, computer took 2, so 5 sticks remain
 Round 1, 5 sticks at start, human took 1, so 4 sticks remain
 Round 2, 4 sticks at start, computer took 2, so 2 sticks remain
 Round 3, 2 sticks at start, human took 1, so 1 sticks remain
 Round 4, 1 sticks at start, computer took 1, so 0 sticks remain
 Computer wins
```

To run program

1. Enter the heap size for the game

```
go run main.go
```
