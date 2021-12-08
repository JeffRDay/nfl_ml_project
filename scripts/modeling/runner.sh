#!/bin/bash

START=$(date)

for i in {1000,500,250,50,25,1}
do
#   echo "Number: $i"
python3 forest.py --limiter $i -e 150
python3 knn.py --limiter $i -k 4
python3 mlr.py --limiter $i
python3 tree.py --limiter $i
python3 mlp.py --limiter $i
done

END=$(date)

echo "START TIME: $START, END TIME: $END"