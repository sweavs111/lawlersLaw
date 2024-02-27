package analysis

import (
	"overUnderModel/game"
	"overUnderModel/numbers"
	"overUnderModel/numbers/logspace"
	"sort"
)

type Point struct {
	WinningScore int     // was this the winning score?
	GamesPlayed  int     // how many games reached or surpassed this score
	GamesWon     int     // how many times was the winning team ahead at this point value
	WinPerc      float64 // what was the winning percentage of the team leading at this score
	PVal         float64 // p-value of games won compared to null hypothesis binomial distribution
}

// WinningScoreDist takes in a []game.Game and returns a map with winning score as the key and frequency as the value.
// It additionally returns the average winning score for that season
func WinningScoreDist(g []game.Game) (map[int]*Point, float64) {
	var n, d float64
	var found bool
	//var p Point
	var j int
	mp := make(map[int]*Point)
	for i := range g {
		if g[i].PointsLed[0] == 0 {
			continue
		}
		_, found = mp[g[i].WinningScore]
		if !found {
			mp[g[i].WinningScore] = &Point{WinningScore: 0, GamesPlayed: 0, GamesWon: 0, WinPerc: 0, PVal: 0}
		}
		mp[g[i].WinningScore].WinningScore++
		n += float64(g[i].WinningScore)
		d++
		for _, j = range g[i].PointsLed {
			_, found = mp[j]
			if !found {
				mp[j] = &Point{WinningScore: 0, GamesPlayed: 0, GamesWon: 0, WinPerc: 0, PVal: 0}
			}
			mp[j].GamesWon++
			//p = mp[g[i].PointsLed[j]]
			//p.GamesWon++
			//mp[g[i].PointsLed[j]] = p
		}
	}
	return mp, n / d
}

func PointsLedDistribution(g []game.Game) map[int]int {
	var j int
	mp := make(map[int]int)
	for i := range g {
		if g[i].PointsLed[0] == 0 {
			continue
		}
		for _, j = range g[i].PointsLed {
			mp[j]++
		}
	}
	return mp
}

func GetDenominator(mp map[int]*Point) {
	//var p Point
	var found bool
	mx, c := mapStats(mp)
	for i := 1; i <= mx; i++ {
		_, found = mp[i]
		if !found {
			mp[i] = &Point{WinningScore: 0, GamesPlayed: 0, GamesWon: 0, WinPerc: 0, PVal: 0}
		}
		mp[i].GamesPlayed = c
		c -= mp[i].WinningScore
	}
}

func mapStats(mp map[int]*Point) (int, int) {
	var mx, c int
	for i := range mp {
		mx = numbers.MaxInt(mx, i)
		//fmt.Println(mp[i].GamesWon)
		c += mp[i].WinningScore
	}
	return mx, c
}

func PointsLedWinPercentage(mp map[int]*Point) {
	for i := range mp {
		mp[i].WinPerc = float64(mp[i].GamesWon) / float64(mp[i].GamesPlayed)
	}
}

func CalcP(mp map[int]*Point) {
	var pts []int
	var p float64
	for i := range mp {
		pts = append(pts, i)
	}
	sort.Ints(pts)
	for i := range pts {
		p = 0
		for k := mp[pts[i]].GamesWon; k <= mp[pts[i]].GamesPlayed; k++ {
			//fmt.Println(k)
			p = logspace.Add(numbers.BinomialDistLog(mp[pts[i]].GamesPlayed, k, 0.5), p)
		}
		mp[pts[i]].PVal = logspace.LogToNormalSpace(p) - 1
	}
}
