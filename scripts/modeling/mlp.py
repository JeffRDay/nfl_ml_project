'''
Creates a Multilayer Perceptron Classifier (MLPC) model. The
model is trained using the csv data provided. Evaulation is
conducted using 5-fold cross validation. The model is serialized
and saved at the end of execution regardless of performance.
'''
import argparse
import csv
import logging
import math
import time
from datetime import datetime

import numpy as np
import pandas as pd
from sklearn.exceptions import UndefinedMetricWarning
from sklearn.neural_network import MLPRegressor
from sklearn.utils._testing import ignore_warnings
from sklearn import preprocessing, model_selection

parser = argparse.ArgumentParser()
parser.add_argument("-l", "--limiter", dest = "LIMITER", type = int, default = 1000, help="rate limiter ::: len(data_size) / limiter")
parser.add_argument("-p", "--path",dest = "PATH", type = str, default = "../../dist/20210712_214459_formatted-data.csv", help="path to csv data file")
parser.add_argument("-layers", "--layers", dest = "LAYERS", type = int, default = 20, help="number of layers in MLP Regressor")
parser.add_argument("-size", "--layer-size", dest = "LAYER_SIZE", type = int, default = 60, help="perceptrons per layer in MLP Regressor")
args = parser.parse_args()

def main():
    ''' main execution function '''
    logging.basicConfig(
        level=logging.INFO,
        format="%(levelname)s ::: %(message)s")

    logging.info("starting mlp.py")
    logging.info("params LIMITER=%s, PATH=%s, LAYERS=%s, LAYER_SIZE=%s", args.LIMITER, args.PATH, args.LAYERS, args.LAYER_SIZE)
    start = time.time()

    data_frame = read_csv(args.PATH)
    logging.info(
        "read in %s records to data frame with %s features",
        len(data_frame),
        len(data_frame.columns))

    size = len(data_frame) // args.LIMITER
    logging.info("will use %s records for MLP Regressor", size)

    subset = data_frame.iloc[:size,]
    working_subset = clean_copy(subset)
    predictors, target = split_data(working_subset)

    logging.info("converting predictors and target to numpy arrays")
    predictors = predictors.to_numpy()
    target = target.to_numpy()

    logging.warning("starting MLP Regressor - good luck.")
    results = model(
        "MLP Regression using Relu: " +str(args.LAYERS)+" hidden layers with "+str(args.LAYER_SIZE)+
        " nodes per layer",
        (args.LAYERS,args.LAYER_SIZE), "relu", predictors, target)

    logging.info("MLP Regressor completed, results:")
    logging.info(results)

    logging.info("updating reports.csv")
    save("../../runs/report.csv", results)

    end = time.time()
    print("executed mlp.py in :", end-start)

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

def model(description, arch, activation, predictors, target):
    ''' uses 5-fold cross validation to create and save MLPC model '''
    kfold = model_selection.KFold(5, shuffle=True, random_state=2)

    rmse, mse, r2 = [], [], []

    for train_idx, test_idx in kfold.split(predictors):
        predictors_train, predictors_test = predictors[train_idx], predictors[test_idx]
        target_train, target_test = target[train_idx], target[test_idx]

        mlp_clf = MLPRegressor(
            hidden_layer_sizes=arch,
            max_iter=2000,
            activation=activation,
            random_state=2,
            solver='lbfgs')

        with ignore_warnings(category=UndefinedMetricWarning):

            predictors_scaler = preprocessing.MinMaxScaler()

            predictors_train = predictors_scaler.fit_transform(predictors_train)

            mlp_clf.fit(predictors_train, target_train)

            predictors_test = predictors_scaler.transform(predictors_test)
            target_prediction = mlp_clf.predict(predictors_test)

            rmse += [math.sqrt((np.mean(target_prediction - target_test) ** 2))]
            mse += [np.mean((target_prediction - target_test) ** 2)]
            r2 += [mlp_clf.score(predictors_test, target_test)]

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
