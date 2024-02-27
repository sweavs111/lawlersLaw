package scrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"overUnderModel/numbers"
	"strconv"
	"strings"
)

func GetLinks(url string) []string {
	var err error
	var links []string
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnHTML("#stats", func(h *colly.HTMLElement) {
		h.ForEachWithBreak("th[data-stat=season]", func(_ int, h *colly.HTMLElement) bool {
			if h.ChildText("a") == "1995-96" {
				return false
			}
			//time.Sleep(5 * time.Second)
			if h.ChildText("a") != "" {
				links = append(links, h.Request.AbsoluteURL(h.ChildAttr("a", "href")))
			}
			return true
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("Visting %s\n", r.URL))
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	err = c.Visit(url)
	if err != nil {
		err.Error()
	}
	return links
}

func NavigatePBP(url string) []string {
	var err error
	var b string
	var pbp []string

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnHTML("#schedule", func(h *colly.HTMLElement) {
		h.ForEachWithBreak("td[data-stat=box_score_text]", func(_ int, h *colly.HTMLElement) bool {
			b = h.ChildAttr("a", "href")
			if len(b) == 0 {
				return false
			}
			pbp = append(pbp, h.Request.AbsoluteURL(b[0:11]+"pbp/"+b[11:]))
			return true
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("Visting %s\n", r.URL))
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	err = c.Visit(url)
	if err != nil {
		err.Error()
	}
	return pbp
}

func GetMonths(url string) []string {
	var err error
	var months []string
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnHTML(".filter", func(h *colly.HTMLElement) {
		h.ForEach("div", func(_ int, h *colly.HTMLElement) {
			months = append(months, h.Request.AbsoluteURL(h.ChildAttr("a", "href")))
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("Visting %s\n", r.URL))
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

	err = c.Visit(url)
	if err != nil {
		err.Error()
	}
	return months
}

func ScrapeBBallRef(url string) (int, int, []int) {
	var score []string
	var data, finalScores []int
	var curr, prev, final, w, l, ws, ls int
	var lead bool = false

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnHTML(".box", func(h *colly.HTMLElement) {
		h.ForEach(".scores", func(_ int, el *colly.HTMLElement) {
			final, _ = strconv.Atoi(el.ChildText(".score"))
			finalScores = append(finalScores, final)
		})

		if finalScores[0] > finalScores[1] {
			w = 0
			l = 1
		} else {
			w = 1
			l = 0
		}

		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			score = strings.Split(el.ChildText("td:nth-child(4)"), "-")
			if len(score) == 2 {
				ws, _ = strconv.Atoi(score[w])
				ls, _ = strconv.Atoi(score[l])
				if ws > ls {
					lead = true
					curr = ws
				} else if lead {
					for i := prev + 1; i <= curr; i++ {
						data = append(data, i)
						lead = false
						prev = ls
					}
				} else {
					prev = ls
				}
			}
		})
		for i := prev + 1; i <= curr; i++ {
			data = append(data, i)
		}
	})

	err := c.Visit(url)
	if err != nil {
		err.Error()
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf(fmt.Sprintf("Visting %s\n", r.URL))
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping %s\n", e.Error())
	})

	if len(finalScores) != 2 {
		fmt.Println(url)
	}

	return numbers.MaxInt(finalScores[0], finalScores[1]), numbers.AbsInt(finalScores[0] - finalScores[1]), data
}
