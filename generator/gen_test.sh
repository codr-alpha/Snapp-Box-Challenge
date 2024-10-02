#!/bin/bash
i=1
while [ $i -le 30 ]
do
	echo "$i 100000" > tmp
	go run main.go < tmp
	((i++))
done
rm tmp