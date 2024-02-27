package rw

import (
	"fmt"
	"overUnderModel/analysis"
	"overUnderModel/convert"
	"overUnderModel/exception"
	"overUnderModel/fileio"
	"overUnderModel/game"
	"sort"
)

func Write(filename *fileio.EasyWriter, g game.Game) {
	pts := convert.SliceIntToString(g.PointsLed)
	s := fmt.Sprintf("%d\t%d\t%d\t%s", g.WinningScore, g.Margin, g.Year, pts)
	fileio.WriteToFileHandle(filename, s)
}

func WriteAvgWinScore(filename string, data []string) {
	fileio.Write(filename, data)
}

func WriteWinPercAt100(filename string, data []string) {
	fileio.Write(filename, data)
}

func WriteWinningScoreDist(mp map[int]*analysis.Point, year int, filename string) {
	var slc []int

	o := fileio.EasyCreate(filename)
	fileio.WriteToFileHandle(o, "score\tfrequency\tyear")
	for i := range mp {
		slc = append(slc, i)
	}
	sort.Ints(slc)
	for i := range slc {
		fileio.WriteToFileHandle(o, fmt.Sprintf("%d\t%d\t%d", slc[i], mp[slc[i]].WinningScore, year))
	}
	err := o.Close()
	exception.PanicOnErr(err)
}

func WritePointsLedDist(mp map[int]*analysis.Point, filename string) {
	var slc []int

	o := fileio.EasyCreate(filename)
	fileio.WriteToFileHandle(o, "score\tgamesWon")
	for i := range mp {
		slc = append(slc, i)
	}
	sort.Ints(slc)
	for i := range slc {
		fileio.WriteToFileHandle(o, fmt.Sprintf("%d\t%d", slc[i], mp[slc[i]].GamesWon))
	}
	err := o.Close()
	exception.PanicOnErr(err)
}

func WritePointsLedWinPercentage(mp map[int]*analysis.Point, filename string) {
	var slc []int

	o := fileio.EasyCreate(filename)
	fileio.WriteToFileHandle(o, "score\twinPercentage\tnGames\tpVal")
	for i := range mp {
		slc = append(slc, i)
	}
	sort.Ints(slc)
	for i := range slc {
		fileio.WriteToFileHandle(o, fmt.Sprintf("%d\t%f\t%d\t%.12f", slc[i], mp[slc[i]].WinPerc, mp[slc[i]].GamesPlayed, mp[slc[i]].PVal))
	}
	err := o.Close()
	exception.PanicOnErr(err)
}
