Yahtzee is a repository I'm working on to practice Docker, Go, and K8s. 

My goal is to make a cluster which contains a few services:
* 'player' services which implement a yahtzee api:
    * Take a turn when asked
        * Call a 'roll' service
        * Decide on rerolls, what trick to use the roll for
        * Return final roll and trick
* 'host' service which coordinates yahtzee games
    * makes calls to registered 'player' services
    * stores results of game
    * API to start games and view results
* 'ui' service to coordinate with the host and display results on a web page
* 'roll' service to generate dice rolls
* Some persistent storage layer for recording game results

Different types of players can be implented and registered.
Then, an end user can start a game with selected players and see the results.
The user can also run multiple simulated games and compare stats for each player's final score.
