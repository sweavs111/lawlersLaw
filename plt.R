library(tidyverse)
library(ggpubr)
library(ggridges)

df = read.delim("~/overUnderModel/data/process/2023.winningScoreDist.txt", sep = "\t")
ggbarplot(df, x = "score", y = "count", fill = 'dark green') +
  xlim(0, 180) + geom_hline(yintercept = 0)

df2008 = read.delim("~/overUnderModel/data/process/2008.winningScoreDist.txt", sep = "\t")
ggbarplot(df2008, x = "score", y = "count", fill = 'dark green') +
  xlim(0, 180) + geom_hline(yintercept = 0)

ggplot(df, aes(score, count)) + geom_bar(fill = 'blue', stat = "identity") + 
  geom_bar(data = df2008, fill = 'red', stat = "identity") + theme_classic()

#some geom_ridge stuff
rdg = read.delim("~/overUnderModel/data/process/allYearsWinningScoreDist.txt", sep ="\t")
rdg$year = as.factor(rdg$year)
ggplot(rdg, aes(x = score, y = year, fill = year)) + geom_density_ridges_gradient(scale = 4, rel_min_height = 0.01) + 
  scale_fill_viridis_d(name = "year", option = "C") + #theme_classic() +
  xlim(0, 180) + geom_vline(xintercept = 102.3, color = "white", linetype = "dashed") + 
  geom_vline(xintercept = 119.9, color = "white", linetype = "dashed")

rdg2 = read.delim("~/overUnderModel/data/process/06to23.winningScoreDist.txt", sep = "\t")
rdg2$year = as.factor(rdg2$year)
ggplot(rdg2, aes(x = score, y = year, fill = year)) + geom_density_ridges(scale = 3, rel_min_height = 0.01) + theme_classic()

####

df2 = read.delim("~/overUnderModel/data/process/2023.pointsLedDist.txt", sep = "\t")
ggbarplot(df2, x = "score", y = "gamesWon", fill = 'dark green') +
  xlim(0, 180) + geom_hline(yintercept = 0)

df3 = read.delim("~/overUnderModel/data/process/2023.pointsLedWinPerc.txt", sep = "\t")
ggbarplot(df3, x = "score", y = "winPercentage", fill = 'pVal', color = 'pVal') +
  geom_hline(yintercept = 1) + geom_hline(yintercept = 0) + scale_fill_gradient(low = "light blue", high = "dark blue") +
  scale_color_gradient(low = "light blue", high = "dark blue")
 
df4 = read.delim("~/overUnderModel/data/process/winningPercAt100.txt", sep = "\t")
ggbarplot(df4, x = "year", y = "winningPerc", fill = "gamesPlayed", color = "gamesPlayed") + coord_cartesian(ylim=c(0.8,1))

#binomal distribution
bd = read.delim("~/overUnderModel/data/process/bd.txt", sep ="\t")
pt = subset(bd, k == 690)
ggplot(bd, aes(x = k, y = prob)) + geom_point() + geom_point(data = pt, color = 'green') + 
  geom_hline(yintercept = 0.05, color = 'red') + geom_hline(yintercept = 0.95, color = 'red') + 
  xlim(575, 725) + theme_classic()

