# GameMode
Normal game mode command.

# How to use
```go
package main

import "github.com/daft0175/gamemode"

func main() {
cmd.Register(cmd.New("gamemode", "Sets a player's game mode.", []string{""}, gamemode.GameMode{}))
}
```
