import csv
import os
import pandas as pd

print("Running check-columsn.py")

# Get all column names in all files
column_names = []
for f in os.listdir("../prototype-data"):
    with open("../prototype-data/" + f) as file:
        csvFile = csv.reader(file)
        for lines in csvFile:
            for column in lines:
                column_names.append(column)
            break

# Print the number of unique column names
unique_column_names = list(dict.fromkeys(column_names))
print("Number of unique column names: " + str(len(unique_column_names)))
print("Column names: " + str(unique_column_names))

dataframes = []
for filename in os.listdir("../data"):
    df = pd.read_csv("../data/" + filename, index_col=None, header=0)
    dataframes.append(df)

count = 0
for df in dataframes:
    count += len(df)
    
print(count)

