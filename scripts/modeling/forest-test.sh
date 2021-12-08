#!/bin/bash

START=$(date)

for i in {25,50,100,150,200,250,300,350}
do
  echo "Number: $i"
python3 forest.py -e $i --limiter 50
done

END=$(date)

echo "START TIME: $START, END TIME: $END"