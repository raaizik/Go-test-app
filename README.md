# Go-test-app
This is a Chess CLI app
## Prerequisites
1. Install task `brew install go-task/tap/go-task`

## Build and run
1. `task build`
2. `task run` to list the available commands and their descriptions
3. Example: 

```
./bin/Go-test-app readFEN "8/p1b5/1pk2p2/2pp2p1/P2P2Pp/1PP1NP2/2KN2n1/8 w - - 0 1"
```
```
./bin/Go-test-app getRandGame
```
```
./bin/Go-test-app findOpening adams
```