package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type GameMode struct {
	Number int
}

func (gm GameMode) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	error := "Syntax error: Unexpected \"\": at \"/gamemode>><<\""
	switch gm.Number {
	case 0:
		p.SetGameMode(world.GameModeSurvival)
		output.Print("Your game mode has been updated to to Survival")
	case 1:
		p.SetGameMode(world.GameModeCreative)
		output.Print("Your game mode has been updated to to Creative")
	case 2:
		p.SetGameMode(world.GameModeAdventure)
		output.Print("Your game mode has been updated to to  Adventure")
	case 3:
		p.SetGameMode(world.GameModeSpectator)
		output.Print("Your game mode has been updated to to Spectator")
	default:
		output.Print(error)
	}
}
