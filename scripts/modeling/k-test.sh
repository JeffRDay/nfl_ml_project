#!/bin/bash

START=$(date)

for i in {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}
do
  echo "Number: $i"
python3 knn.py -k $i --limiter 50
done

END=$(date)

echo "START TIME: $START, END TIME: $END"