#!/bin/bash

# python3 forest.py
# python3 knn.py
# python3 mlp.py
# python3 mlr.py
# python3 tree.py

START=$(date)

for i in {1000,2000}
do
#   echo "Number: $i"
python3 forest.py --limiter $i
python3 knn.py --limiter $i -k 4
python3 mlp.py --limiter $i
python3 mlr.py --limiter $i
python3 tree.py --limiter $i
done

END=$(date)

echo "START TIME: $START, END TIME: $END"