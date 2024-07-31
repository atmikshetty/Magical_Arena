# Magical Arena Game

Magical Arena is a Go-based game simulation where players with different attributes battle in an arena using dice rolls to determine the outcome of their attacks and defenses.

# File Structure of the Application

Game(root Directory)
    arena:
        arena_test.go
        arena.go
    dice:
        dice_test.go
        dice.go
    player:
        player_test.go
        player.go
    go.mod
    main.go
ReadMe.md

# The code is explained in detailed in every single go file in the form of comments

## Steps to Run the Application

1. Ensure you have Go installed on your system. If not, visit the official GO Documentation and install it for your Operating System.

2. Extract everything from the Zip file 

2. Move to this directory: cd game

3. Run the application: go run main.go

This will start the Game Arena simulation, creating two players and simulating a battle between them. The output will be continously printed to your terminal. 

## Running Tests

To run all tests for the project:

1. Navigate to the the game directory. 
2. Run the following command: go test ./... 


This command runs all tests in all subdirectories and displays the output accordingly. All the tests will be passed. 

To run tests for a specific package:

1. Navigate to the the game directory.
2. Run: go test

# Testing with Detailed Output
To run tests with verbose output: go test -v ./...

To run a specific test function: go test -v -run TestFunctionName

Replace `TestFunctionName` with the name of the test function you want to run.

## Accessing and Writing Tests

Tests are located in `*_test.go` files within each package directory. To add or modify tests:

1. Open the relevant `*_test.go` file (e.g., `dice/dice_test.go` for dice tests).
2. Add new test functions or modify existing ones. Test function names should start with `Test`.
3. Use `t.Errorf()` to report test failures.

Example test function:

```go
func TestDiceRoll(t *testing.T) {
    d := NewDice()
    roll := d.Roll()
    if roll < 1 || roll > 6 {
        t.Errorf("Roll %d is out of range [1,6]", roll)
    }
}
```

# Git History

As mentioned in the Google Form, local git commits have been made frequently. 
Use 'git reflog' command to see all the commits.

