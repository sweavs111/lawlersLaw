#!/bin/bash

for i in *pointsLedWinPerc.txt
do
	echo $i
	cat $i | awk '{if ($1 == 100) print $0}'
done
