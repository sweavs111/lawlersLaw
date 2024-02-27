package main

import (
	"flag"
	"fmt"
	"log"
	"overUnderModel/analysis"
	"overUnderModel/exception"
	"overUnderModel/fileio"
	"overUnderModel/game"
	"overUnderModel/rw"
	"overUnderModel/scrape"
	"strconv"
	"time"
)

func analyze() {
	var g []game.Game
	var avg float64 // avg winning score
	var mp map[int]*analysis.Point
	var winningScoreAvg []string = []string{"year\tavgScore"}
	var winPerAt100 []string = []string{"year\twinPercentage"}
	for i := 1997; i < 2025; i++ {
		g = rw.ReadGames(fmt.Sprintf("data/raw/%d.pbp.txt", i))
		mp, avg = analysis.WinningScoreDist(g)
		winningScoreAvg = append(winningScoreAvg, fmt.Sprintf("%d\t%f", i, avg))
		rw.WriteWinningScoreDist(mp, i, fmt.Sprintf("data/process/%d.winningScoreDist.txt", i))
		//mp2 = analysis.PointsLedDistribution(g)
		rw.WritePointsLedDist(mp, fmt.Sprintf("data/process/%d.pointsLedDist.txt", i))
		analysis.GetDenominator(mp)
		analysis.PointsLedWinPercentage(mp)
		winPerAt100 = append(winPerAt100, fmt.Sprintf("%d\t%f", i, mp[100].WinPerc))
		analysis.CalcP(mp)
		rw.WritePointsLedWinPercentage(mp, fmt.Sprintf("data/process/%d.pointsLedWinPerc.txt", i))
	}
	rw.WriteWinPercAt100("data/process/winPercentAt100ByYear.txt", winPerAt100)
	rw.WriteAvgWinScore("data/process/avgWinningScoreByYear.txt", winningScoreAvg)
}

/*
https://www.basketball-reference.com/leagues/NBA_2024.html  0  up to date with 2/21/24
https://www.basketball-reference.com/leagues/NBA_2023.html  1  done
https://www.basketball-reference.com/leagues/NBA_2022.html  2  done
https://www.basketball-reference.com/leagues/NBA_2021.html  3  done
https://www.basketball-reference.com/leagues/NBA_2020.html  4  done
https://www.basketball-reference.com/leagues/NBA_2019.html  5  done
https://www.basketball-reference.com/leagues/NBA_2018.html  6  done
https://www.basketball-reference.com/leagues/NBA_2017.html  7  done
https://www.basketball-reference.com/leagues/NBA_2016.html  8  done
https://www.basketball-reference.com/leagues/NBA_2015.html  9  done
https://www.basketball-reference.com/leagues/NBA_2014.html  10  done
https://www.basketball-reference.com/leagues/NBA_2013.html  11  done
https://www.basketball-reference.com/leagues/NBA_2012.html  12  done
https://www.basketball-reference.com/leagues/NBA_2011.html  13  done
https://www.basketball-reference.com/leagues/NBA_2010.html  14  done
https://www.basketball-reference.com/leagues/NBA_2009.html  15  done
https://www.basketball-reference.com/leagues/NBA_2008.html  16  done
https://www.basketball-reference.com/leagues/NBA_2007.html  17  done
https://www.basketball-reference.com/leagues/NBA_2006.html  18  done
https://www.basketball-reference.com/leagues/NBA_2005.html  19  done
https://www.basketball-reference.com/leagues/NBA_2004.html  20  done
https://www.basketball-reference.com/leagues/NBA_2003.html  21  done
https://www.basketball-reference.com/leagues/NBA_2002.html  22  done
https://www.basketball-reference.com/leagues/NBA_2001.html  23  done
https://www.basketball-reference.com/leagues/NBA_2000.html  24  done
https://www.basketball-reference.com/leagues/NBA_1999.html  25  done
https://www.basketball-reference.com/leagues/NBA_1998.html  26  done
https://www.basketball-reference.com/leagues/NBA_1997.html  27  done
*/

func getData() {
	var g game.Game
	var strip string
	var gameLinks, pbp []string

	idx := 0

	seasons := scrape.GetLinks("https://www.basketball-reference.com/leagues/")

	for i := range seasons {
		strip = seasons[i][0 : len(seasons[i])-5]
		gameLinks = append(gameLinks, strip+"_games.html")
	}

	time.Sleep(4 * time.Second)

	months := scrape.GetMonths(gameLinks[idx])
	g.Year, _ = strconv.Atoi(gameLinks[idx][len(gameLinks[1])-15 : len(gameLinks[5])-11])
	o := fileio.EasyCreate(fmt.Sprintf("data/raw/%d.1.pbp.txt", g.Year))

	for i := 3; i < len(months); i++ {
		fileio.WriteToFileHandle(o, "#"+months[i])
		time.Sleep(4 * time.Second)
		pbp = scrape.NavigatePBP(months[i])
		for j := range pbp {
			time.Sleep(3*time.Second + 500*time.Millisecond)
			g.WinningScore, g.Margin, g.PointsLed = scrape.ScrapeBBallRef(pbp[j])
			rw.Write(o, g)
		}
	}
	err := o.Close()
	exception.PanicOnErr(err)
}

func main() {
	flag.Parse()
	if flag.Arg(0) == "s" {
		getData()
	} else if flag.Arg(0) == "a" {
		analyze()
	} else {
		log.Fatalf("option must either be 's' (scrape) or 'a' (analyze)")
	}
}
