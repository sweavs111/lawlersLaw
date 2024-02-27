package rw

import (
	"fmt"
	"overUnderModel/convert"
	"overUnderModel/fileio"
	"overUnderModel/game"
	"strings"
)

func ReadGames(file string) []game.Game {
	var g []game.Game
	slc := fileio.Read(file)

	for i := range slc {
		if slc[i][0] == '#' {
			fmt.Println(slc[i])
			continue
		}
		g = append(g, lineToGame(slc[i]))
	}
	return g
}

func lineToGame(s string) game.Game {
	var gm game.Game
	col := strings.Split(s, "\t")
	gm.WinningScore = convert.StringToInt(col[0])
	gm.Margin = convert.StringToInt(col[1])
	gm.Year = convert.StringToInt(col[2])
	gm.PointsLed = ptsStringToSlice(col[3])
	return gm
}

func ptsStringToSlice(s string) []int {
	var slc []int
	col := strings.Split(s, ",")
	for i := range col {
		slc = append(slc, convert.StringToInt(col[i]))
	}
	return slc
}
