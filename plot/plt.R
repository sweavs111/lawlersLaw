library(tidyverse)
library(ggpubr)
library(ggridges)

# plot winning score distribution
for (x in 1997:2024) {
  str = paste("~/lawlersLaw/data/process/", as.character(x), ".winningScoreDist.txt", sep = "")
  df = read.delim(str, sep = "\t")
  p = ggbarplot(df, x = "score", y = "frequency", fill = 'dark green', color = 'dark green') +
    xlim(0, max(df$score+10)) + geom_hline(yintercept = 0) + xlab("Winning score") + ylab("Frequency") +
    ggtitle(paste(x, "winning score distribution"))
  ggsave(paste("~/lawlersLaw/plot/winningScoreDistByYear/", as.character(x), ".winningScoreDist.png", sep = ""), plot = p)
}

# plot points led distribution
for (x in 1997:2024) {
  str = paste("~/lawlersLaw/data/process/", as.character(x), ".pointsLedDist.txt", sep = "")
  df = read.delim(str, sep = "\t")
  p = ggbarplot(df, x = "score", y = "gamesWon", fill = 'light blue', color = 'light blue') +
    xlim(0, max(df2$score)) + geom_hline(yintercept = 0) + ylab("Number of games won") +
    xlab("Score of team in the lead") + ggtitle(paste(as.character(x), "number games won by team in the lead at a given score"))
  ggsave(paste("~/lawlersLaw/plot/pointsLedDistByYear/", as.character(x), ".pointsLedDist.png", sep = ""), plot = p)
}

# plot points led win percentage
for (x in 1997:2024) {
  str = paste("~/lawlersLaw/data/process/", as.character(x), ".pointsLedWinPerc.txt", sep = "")
  df = read.delim(str, sep = "\t")
  p = ggbarplot(df, x = "score", y = "winPercentage", fill = 'orange', color = 'orange') +
    geom_hline(yintercept = 1) + geom_hline(yintercept = 0) + xlab("Score of team in the lead") +
    ylab("Winning percentage") + ggtitle(paste(as.character(x), "winning percentage of team in the lead at a given score"))
  ggsave(paste("~/lawlersLaw/plot/pointsLedWinPercByYear/", as.character(x), ".pointsLedWinPerc.png", sep = ""), plot = p)
}




# geom_ridge winning score dist by year
rdg = read.delim("~/lawlersLaw/data/process/allYearsWinningScoreDist.txt", sep ="\t")
rdg$year = as.factor(rdg$year)
p = ggplot(rdg, aes(x = score, y = year, fill = year)) + geom_density_ridges_gradient(scale = 3, rel_min_height = 0.01, bandwidth = 6) + 
  scale_fill_viridis_d(name = "year", option = "C") + theme_classic() + theme(legend.position = "none") +
  xlim(0, 180) + xlab("Winning Score Distribution") + ggtitle("Winning score distribution by year")
ggsave("~/lawlersLaw/plot/winningScoreDistByYear.ridge.png", plot = p)
# avg winning score by year
avgScore = read.delim("~/lawlersLaw/data/process/avgWinningScoreByYear.txt", sep = "\t")
avgScore$year = as.factor(avgScore$year)
p = ggbarplot(avgScore, x = "year", y = "avgScore", fill = "avgScore") + coord_cartesian(ylim = c(90, 125)) + 
  ylab("Average Winning Score") + theme(legend.position = "none", axis.text.x = element_text(angle = 45, hjust = 1)) + ggtitle('Average winning score by year')
ggsave("~/lawlersLaw/plot/avgWinningScoreByYear.png", plot = p)

#winning percentage at 100, lawlers law value
df = read.delim("~/lawlersLaw/data/process/winPercentAt100ByYear.txt", sep = "\t")
df$year = as.factor(df$year)
p = ggbarplot(df, x = "year", y = "winPercentage", fill = "gamesPlayed", color = "gamesPlayed") +
  coord_cartesian(ylim=c(0.8,1)) + theme(axis.text.x = element_text(angle = 45, hjust = 1)) +
  ylab("Winning Percentage") + xlab("Year") + ggtitle("Lawlers Law value by year. (Win % of first team to 100)")
ggsave("~/lawlersLaw/plot/lawlersLawValueByYear.png", plot = p)

#binomal distribution
bd = read.delim("~/overUnderModel/data/process/bd.txt", sep ="\t")
pt = subset(bd, k == 690)
ggplot(bd, aes(x = k, y = prob)) + geom_point() + geom_point(data = pt, color = 'green') + 
  geom_hline(yintercept = 0.05, color = 'red') + geom_hline(yintercept = 0.95, color = 'red') + 
  xlim(575, 725) + theme_classic()

