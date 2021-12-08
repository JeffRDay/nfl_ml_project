'''
Creates a K Nearest Neighbor (KNN) model. The model is trained
using the csv data provided. Evaulation is conducted using
5-fold cross validation. The model is serialized and saved at
the end of execution regardless of performance.
'''
import argparse
import csv
import logging
import math
import time
from datetime import datetime

import pandas as pd
import numpy as np
from sklearn import neighbors, model_selection

parser = argparse.ArgumentParser()
parser.add_argument("-l", "--limiter", dest = "LIMITER", type = int, default = 1000, help="rate limiter ::: len(data_size) / limiter")
parser.add_argument("-k", "--k-value", dest = "K_VALUE", type = int, default = 5, help="set k value in knn")
parser.add_argument("-p", "--path",dest = "PATH", type = str, default = "../../dist/20210712_214459_formatted-data.csv", help="path to csv data file")
args = parser.parse_args()

def main():
    ''' main execution function '''
    logging.basicConfig(
        encoding='utf-8',
        level=logging.INFO,
        format="%(levelname)s ::: %(message)s")

    logging.info("starting knn.py")
    logging.info("params LIMITER=%s, K_VALUE=%s, PATH=%s", args.LIMITER, args.K_VALUE, args.PATH)
    start = time.time()

    data_frame = read_csv(args.PATH)
    logging.info(
        "read in %s records to data frame with %s features",
        len(data_frame),
        len(data_frame.columns))

    size = len(data_frame) // args.LIMITER
    logging.info("will use %s records for MLP Classifier", size)

    subset = data_frame.iloc[:size,]
    working_subset = clean_copy(subset)
    predictors, target = split_data(working_subset)

    logging.info("converting predictors and target to numpy arrays")
    predictors = predictors.to_numpy()
    target = target.to_numpy()

    logging.warning("starting KNN Regression - good luck.")
    results = model(
        "KNN Regressor, where K = " +str(args.K_VALUE),
        (args.K_VALUE),predictors, target)

    logging.info("KNN Regression completed, results:")
    logging.info(results)

    logging.info("updating reports.csv")
    save("../../runs/report.csv", results)

    end = time.time()
    print("executed knn.py in :", end-start)

def read_csv(path):
    ''' read_csv reads csv from provided path and return dataframe '''
    return pd.read_csv(path)

def clean_copy(data_frame):
    ''' clean_copy copies the given data frame and drops player_id '''
    wdf = data_frame.copy()
    wdf = wdf.drop('player_id', axis=1)
    return wdf

def split_data(data_frame):
    ''' splits dataframe into predictor and target dataframes '''
    predictors = data_frame.drop('fs_total', axis=1)
    target = data_frame['fs_total']
    return predictors, target

def model(description, K_VALUE, predictors, target):
    ''' uses 5-fold cross validation to create and save MLPC model '''
    kfold = model_selection.KFold(5, shuffle=True, random_state=2)

    rmse, mse, r2 = [], [], []

    for train_idx, test_idx in kfold.split(predictors):
        predictors_train, predictors_test = predictors[train_idx], predictors[test_idx]
        target_train, target_test = target[train_idx], target[test_idx]

        knn_regression = neighbors.KNeighborsRegressor(n_neighbors=K_VALUE, weights='uniform')

        knn_regression.fit(predictors_train, target_train)

        rmse += [math.sqrt(np.mean((knn_regression.predict(predictors_test) - target_test) ** 2))]
        mse += [np.mean((knn_regression.predict(predictors_test) - target_test) ** 2)]
        r2 += [knn_regression.score(predictors_test, target_test)]

    now = datetime.now()
    current_time = now.strftime("%D %H:%M:%S")
    results = [
        description,
        "{:.4f}".format(np.mean(rmse)), # rmse
        "{:.4f}".format(np.std(rmse)), # rmse variance
        "{:.4f}".format(np.mean(mse)), # mse
        "{:.4f}".format(np.std(mse)), # mse variance
        "{:.4f}".format(np.mean(r2)), #r2
        "{:.4f}".format(np.std(r2)), #r2 variance
        "{}".format(len(predictors)), #num records analyzed
        "{}".format(current_time) #time of execution
    ]
    return results

def save(path, contents):
    # Open file in append mode
    with open(path, 'a+', newline='\n') as write_obj:
        # Create a writer object from csv module
        csv_writer = csv.writer(write_obj)
        # Add contents of list as last row in the csv file
        csv_writer.writerow(contents)

main()
